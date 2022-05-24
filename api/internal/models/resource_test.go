package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestResource_NewResource(t *testing.T) {

	models.DbConn = initDB()

	resource := &models.Resource{
		Url: "file:///home/kiritonya/Scaricati/Nana_01_ITA.mp4",
		Language: &models.Language{
			Id:           0,
			Abbreviation: "ita",
			Name:         "italian",
		},
		Quality: &models.Quality{
			Id:         0,
			Quality:    "630p",
			Resolution: "630x480",
		},
		Episode: &models.Episode{
			Series: &models.Series{
				OriginalTitle: &models.Title{
					Title: "NANA",
					Language: &models.Language{
						Abbreviation: "jpn",
						Name:         "japan",
					},
				},
				Plot: "2 ragazze di nome NANA si incontrano",
				Format: &models.Format{
					Type: "Serie",
				},
				Favorites:      100000,
				AgeRestriction: 16,
			},
			Number: 1,
			OriginalTitle: &models.Title{
				Title: "Le due NANA si incontrano",
				Language: &models.Language{
					Abbreviation: "ita",
					Name:         "italian",
				},
			},
			Plot:           "Le due nana si incontrano ed ho sonno",
			AgeRestriction: 16,
		},
	}

	resourceId, err := resource.NewResource()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Resource:", resourceId)
}

func TestGetResourceById(t *testing.T) {

	models.DbConn = initDB()

	resource, err := models.GetResourceById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Resource:", resource)
}
