package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestQuality_NewQuality(t *testing.T) {

	models.DbConn = initDB()

	quality := models.Quality{
		Quality:    "1080p",
		Resolution: "1920x1080",
	}

	qualityId, err := quality.NewQuality()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Id:", qualityId)
}

func TestGetQualityById(t *testing.T) {

	models.DbConn = initDB()

	quality, err := models.GetQualityById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Quality:", quality)
}
