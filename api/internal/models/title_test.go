package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestNewTitle(t *testing.T) {

	models.DbConn = initDB()

	// Create title test
	title := &models.Title{
		Title: "Citrus",
		Language: &models.Language{
			Id:           11,
			Abbreviation: "jpn",
			Name:         "Japanese",
		},
	}

	// Add title
	id, err := title.NewTitle()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ID:", id)
}

func TestGetTitleById(t *testing.T) {

	models.DbConn = initDB()

	title, err := models.GetTitleById(8)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Title:", title)
}
