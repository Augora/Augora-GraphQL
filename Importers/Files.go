package Importers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func getFiles() []Models.DossierHandler {
	filesResp, err := http.Get("https://www.nosdeputes.fr/15/dossiers/date/json")
	if err != nil {
		log.Fatalln(err)
	}

	var dossiers Models.Dossiers
	json.NewDecoder(filesResp.Body).Decode(&dossiers)

	return dossiers.Sections
}

func ImportFiles() {
	db := Utils.GetDataBaseConnection()

	// Loading database models
	db.AutoMigrate(&Models.Dossier{})

	// Begin transation
	tx := db.Begin()

	// Clear current data
	tx.Unscoped().Delete(&Models.Dossier{})

	// Inserting deputes
	for _, file := range getFiles() {
		file.Dossier.ID = 0
		tx.Create(&file.Dossier)
	}

	// Committing transaction
	tx.Commit()
}
