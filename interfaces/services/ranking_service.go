package services

import (
	"net/http"

	"github.com/srvc/fail"
	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/usecases"
)

// RankingService is responsible for managing rankings
type RankingService interface {
	Create(ctx Context) error
}

// NewRankingService initializer
func NewRankingService(rankingInteractor usecases.RankingInteractor) RankingService {
	return &rankingService{
		RankingInteractor: rankingInteractor,
	}
}

type rankingService struct {
	RankingInteractor usecases.RankingInteractor
}

// CreateRankingPayload payload for the create action
type CreateRankingPayload struct {
	ContestID uint64               `json:"contest_id"`
	Languages domain.LanguageCodes `json:"languages"`
}

func (s *rankingService) Create(ctx Context) error {
	payload := &CreateRankingPayload{}
	if err := ctx.Bind(payload); err != nil {
		return fail.Wrap(err)
	}

	user, err := ctx.User()
	if err != nil {
		return fail.Wrap(err)
	}

	if err := s.RankingInteractor.CreateRanking(user.ID, payload.ContestID, payload.Languages); err != nil {
		return fail.Wrap(err)
	}

	return ctx.NoContent(http.StatusCreated)
}