package Importers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
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
	user := os.Getenv("backend_sql_user")
	pass := os.Getenv("backend_sql_password")
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora.database.windows.net:1433?database=augora-db")
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)

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
