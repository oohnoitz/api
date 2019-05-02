package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/services"
	"github.com/tadoku/api/usecases"
)

func TestContestLogService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := &domain.ContestLog{
		ContestID: 1,
		Language:  domain.Japanese,
		Amount:    10,
		MediumID:  1,
	}

	ctx := services.NewMockContext(ctrl)
	ctx.EXPECT().NoContent(201)
	ctx.EXPECT().User().Return(&domain.User{ID: 1}, nil)
	ctx.EXPECT().Bind(gomock.Any()).Return(nil).SetArg(0, *log)

	i := usecases.NewMockRankingInteractor(ctrl)
	i.EXPECT().CreateLog(domain.ContestLog{ContestID: 1, UserID: 1, Language: domain.Japanese, Amount: 10, MediumID: 1}).Return(nil)

	s := services.NewContestLogService(i)
	err := s.Create(ctx)

	assert.NoError(t, err)
}