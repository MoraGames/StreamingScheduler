package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestSeries_NewSeries(t *testing.T) {

	models.DbConn = initDB()

	series := &models.Series{
		OriginalTitle: &models.Title{
			Title: "KissXSis",
			Language: &models.Language{
				Abbreviation: "ita",
				Name:         "italian",
			},
		},
		Plot: "Vooooh",
		Format: &models.Format{
			Type: "AhBoh",
		},
		Favorites:      12,
		AgeRestriction: 18,
	}

	seriesId, err := series.NewSeries()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Series:", seriesId)
}
