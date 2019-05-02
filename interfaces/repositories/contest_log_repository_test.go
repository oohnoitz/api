package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/repositories"
)

func TestContestLogRepository_StoreContest(t *testing.T) {
	t.Parallel()
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewContestLogRepository(sqlHandler)
	log := &domain.ContestLog{
		ContestID: 1,
		UserID:    1,
		Language:  domain.Japanese,
		Amount:    10,
		MediumID:  1,
	}

	{
		err := repo.Store(*log)
		assert.NoError(t, err)
	}

	{
		updatedLog := &domain.ContestLog{
			ID:        1,
			ContestID: 1,
			UserID:    1,
			Language:  domain.Japanese,
			Amount:    10,
			MediumID:  1,
		}
		err := repo.Store(*updatedLog)
		assert.EqualError(t, err, "not yet implemented")
	}
}

func TestContestLogRepository_FindAllByContestAndUser(t *testing.T) {
	t.Parallel()
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewContestLogRepository(sqlHandler)

	contestID := uint64(1)
	userID := uint64(1)

	expected := []struct {
		language domain.LanguageCode
		medium   domain.MediumID
		amount   float32
	}{
		{domain.Japanese, domain.MediumBook, 10},
		{domain.Korean, domain.MediumManga, 20},
		{domain.Global, domain.MediumNet, 30},
	}

	// Correct logs
	{
		for _, data := range expected {
			log := &domain.ContestLog{
				ContestID: contestID,
				UserID:    userID,
				Language:  data.language,
				MediumID:  data.medium,
				Amount:    data.amount,
			}

			err := repo.Store(*log)
			assert.NoError(t, err)
		}
	}

	// Create unrelated rankings to check if it is really working
	{
		for _, language := range []domain.LanguageCode{domain.Korean, domain.Global} {
			ranking := &domain.ContestLog{
				ContestID: contestID + 1,
				UserID:    userID,
				Language:  language,
				MediumID:  domain.MediumBook,
				Amount:    0,
			}

			err := repo.Store(*ranking)
			assert.NoError(t, err)
		}
	}

	logs, err := repo.FindAll(contestID, userID)
	assert.NoError(t, err)

	for _, expected := range expected {
		var log domain.ContestLog
		for _, l := range logs {
			if l.Language == expected.language {
				log = l
			}
		}

		assert.Equal(t, expected.amount, log.Amount)
		assert.Equal(t, expected.medium, log.MediumID)
		assert.Equal(t, contestID, log.ContestID)
		assert.Equal(t, userID, log.UserID)
	}
}