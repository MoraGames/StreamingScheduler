package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestEpisode_NewEpisode(t *testing.T) {

	models.DbConn = initDB()

	episode := &models.Episode{
		Series: &models.Series{
			OriginalTitle: &models.Title{
				Title: "Yagate Kimi ni naru",
				Language: &models.Language{
					Abbreviation: "jpn",
					Name:         "japan",
				},
			},
			Plot: "Yuriiii",
			Format: &models.Format{
				Type: "OVA",
			},
			Favorites:      5000,
			AgeRestriction: 18,
		},
		Number: 0,
		OriginalTitle: &models.Title{
			Title: "Weela",
			Language: &models.Language{
				Abbreviation: "jpn",
				Name:         "japanese",
			},
		},
		Plot:           "Non so un cabbo",
		AgeRestriction: 18,
	}

	episodeId, err := episode.NewEpisode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Id:", episodeId)
}

func TestGetEpisodeById(t *testing.T) {

	models.DbConn = initDB()

	episode, err := models.GetEpisodeById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Episode:", episode)
}
