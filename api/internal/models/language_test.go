package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestLanguage_NewLanguage(t *testing.T) {

	models.DbConn = initDB()

	lang := &models.Language{
		Abbreviation: "jpn",
		Name:         "japanese",
	}

	lastId, err := lang.NewLanguage()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ID:", lastId)
}

func TestGetLanguageById(t *testing.T) {

	models.DbConn = initDB()

	lang, err := models.GetLanguageById(5)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Language:", lang)
}

func TestDeleteLanguage(t *testing.T) {

	models.DbConn = initDB()

	err := models.DeleteLanguage(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("DELETE [OK]")
}
