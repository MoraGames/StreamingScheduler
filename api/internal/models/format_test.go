package models_test

import (
	"github.com/MoraGames/StreamingScheduler/core/internal/models"
	"testing"
)

func TestFormat_NewFormat(t *testing.T) {

	models.DbConn = initDB()

	format := &models.Format{
		Type: "1080p",
	}

	formatId, err := format.NewFormat()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Format:", formatId)
}

func TestGetFormatById(t *testing.T) {

	models.DbConn = initDB()

	format, err := models.GetFormatById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Format:", format)
}

func TestDeleteFormat(t *testing.T) {

	models.DbConn = initDB()

	err := models.DeleteFormat(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("DELETE [OK]")
}
