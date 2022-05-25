package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
	"time"
)

func TestEvent_NewEvent(t *testing.T) {

	models.DbConn = initDB()

	event := models.Event{
		Title:     "Citrus Ep 02 ITA",
		StartTime: time.Now().Add(time.Hour * 25),
		EndTime:   time.Now().Add(time.Hour * 27),
		Resource: &models.Resource{
			Url: "file:///home/kiritonya/Scaricati/Citrus_02_ITA.mp4",
			Language: &models.Language{
				Abbreviation: "ita",
				Name:         "italian",
			},
			Quality: &models.Quality{
				Quality:    "630p",
				Resolution: "630x480",
			},
			Episode: &models.Episode{
				Series: &models.Series{
					OriginalTitle: &models.Title{
						Title: "Citrus",
						Language: &models.Language{
							Abbreviation: "jpn",
							Name:         "japanese",
						},
					},
					Format: &models.Format{
						Type: "SERIE",
					},
					Favorites:      12555,
					AgeRestriction: 16,
				},
				Number: 1,
				OriginalTitle: &models.Title{
					Title: "Le due sorelle si incontrano",
					Language: &models.Language{
						Abbreviation: "ita",
						Name:         "italian",
					},
				},
				AgeRestriction: 16,
			},
		},
	}

	eventId, err := event.NewEvent()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Id:", eventId)
}

func TestGetEventById(t *testing.T) {

	models.DbConn = initDB()

	event, err := models.GetEventById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Event:", event)
}
