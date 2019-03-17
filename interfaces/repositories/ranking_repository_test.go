package repositories_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/repositories"
)

func TestRankingRepository_StoreRanking(t *testing.T) {
	t.Parallel()
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewRankingRepository(sqlHandler)
	ranking := &domain.Ranking{
		ID:        1,
		ContestID: 1,
		UserID:    1,
		Language:  domain.Japanese,
		Amount:    0,
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	{
		err := repo.Store(*ranking)
		assert.Nil(t, err)
	}

	{
		updatedRanking := &domain.Ranking{
			ID:     1,
			Amount: 2,
		}
		err := repo.Store(*updatedRanking)
		assert.Nil(t, err)
	}
}

func TestRankingRepository_GetAllLanguagesForContestAndUser(t *testing.T) {
	t.Parallel()
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewRankingRepository(sqlHandler)
	rankingJapanese := &domain.Ranking{
		ContestID: 1,
		UserID:    1,
		Language:  domain.Japanese,
		Amount:    0,
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	rankingChinese := &domain.Ranking{
		ContestID: 1,
		UserID:    1,
		Language:  domain.Chinese,
		Amount:    0,
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	rankingGlobal := &domain.Ranking{
		ContestID: 1,
		UserID:    1,
		Language:  domain.Global,
		Amount:    0,
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	rankingSingleLanguage := &domain.Ranking{
		ContestID: 1,
		UserID:    2,
		Language:  domain.Chinese,
		Amount:    0,
		CreatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	{
		for _, r := range []*domain.Ranking{rankingJapanese, rankingChinese, rankingGlobal, rankingSingleLanguage} {
			err := repo.Store(*r)
			assert.Nil(t, err)
		}
	}

	{
		languages, err := repo.GetAllLanguagesForContestAndUser(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, len(languages), 3)
		assert.Equal(t, languages[0], domain.Japanese)
		assert.Equal(t, languages[1], domain.Chinese)
		assert.Equal(t, languages[2], domain.Global)
	}

	{
		languages, err := repo.GetAllLanguagesForContestAndUser(1, 2)
		assert.Nil(t, err)
		assert.Equal(t, len(languages), 1)
		assert.Equal(t, languages[0], domain.Chinese)
	}
}