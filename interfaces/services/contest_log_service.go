package services

import (
	"net/http"
	"strconv"

	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/usecases"
)

// ContestLogService is responsible for managing contest log entries
type ContestLogService interface {
	Create(ctx Context) error
	Update(ctx Context) error
	Delete(ctx Context) error
	Get(ctx Context) error
}

// NewContestLogService initializer
func NewContestLogService(rankingInteractor usecases.RankingInteractor) ContestLogService {
	return &contestLogService{
		RankingInteractor: rankingInteractor,
	}
}

type contestLogService struct {
	RankingInteractor usecases.RankingInteractor
}

func (s *contestLogService) Create(ctx Context) error {
	log := &domain.ContestLog{}
	if err := ctx.Bind(log); err != nil {
		return domain.WrapError(err)
	}

	user, err := ctx.User()
	if err != nil {
		return domain.WrapError(err)
	}

	log.UserID = user.ID

	if err := s.RankingInteractor.CreateLog(*log); err != nil {
		return domain.WrapError(err)
	}

	return ctx.NoContent(http.StatusCreated)
}

func (s *contestLogService) Update(ctx Context) error {
	log := &domain.ContestLog{}
	if err := ctx.Bind(log); err != nil {
		return domain.WrapError(err)
	}

	ctx.BindID(&log.ID)

	user, err := ctx.User()
	if err != nil {
		return domain.WrapError(err)
	}

	log.UserID = user.ID

	if err := s.RankingInteractor.UpdateLog(*log); err != nil {
		if err == domain.ErrInsufficientPermissions {
			return ctx.NoContent(http.StatusForbidden)
		}

		return domain.WrapError(err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s *contestLogService) Delete(ctx Context) error {
	var id uint64
	ctx.BindID(&id)

	user, err := ctx.User()
	if err != nil {
		return domain.WrapError(err)
	}

	if err := s.RankingInteractor.DeleteLog(id, user.ID); err != nil {
		if err == domain.ErrInsufficientPermissions {
			return ctx.NoContent(http.StatusForbidden)
		}
		if err == domain.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return domain.WrapError(err)
	}

	return ctx.NoContent(http.StatusOK)
}

func (s *contestLogService) Get(ctx Context) error {
	contestID, err := strconv.ParseUint(ctx.QueryParam("contest_id"), 10, 64)
	if err != nil {
		return domain.WrapError(err)
	}

	userID, err := strconv.ParseUint(ctx.QueryParam("user_id"), 10, 64)
	if err != nil {
		return domain.WrapError(err)
	}

	logs, err := s.RankingInteractor.ContestLogs(contestID, userID)
	if err != nil {
		if err == usecases.ErrNoContestLogsFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return domain.WrapError(err)
	}

	return ctx.JSON(http.StatusOK, logs.GetView())
}
