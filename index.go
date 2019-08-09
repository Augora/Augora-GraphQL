package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KevinBacas/Gin-Go-Test/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func getDeputies() []Models.DeputeHandler {
	deputesResp, err := http.Get("https://www.nosdeputes.fr/deputes/json")
	if err != nil {
		log.Fatalln(err)
	}

	deputesEnMandatResp, err := http.Get("https://www.nosdeputes.fr/deputes/enmandat/json")
	if err != nil {
		log.Fatalln(err)
	}

	var deputes Models.Deputes
	json.NewDecoder(deputesResp.Body).Decode(&deputes)

	var deputesEnMandat Models.Deputes
	json.NewDecoder(deputesEnMandatResp.Body).Decode(&deputesEnMandat)

	for _, deputeEnMandat := range deputesEnMandat.Deputes {
		for deputeIndex, _ := range deputes.Deputes {
			if deputes.Deputes[deputeIndex].Depute.Slug == deputeEnMandat.Depute.Slug {
				deputes.Deputes[deputeIndex].Depute.EstEnMandat = true
			}
		}
	}

	return deputes.Deputes
}

func main() {
	user := os.Getenv("backend_sql_user")
	pass := os.Getenv("backend_sql_password")
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@augora.database.windows.net:1433?database=augora-db")
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)

	// Loading database models
	db.AutoMigrate(&Models.Depute{})
	db.AutoMigrate(&Models.Site{})
	db.AutoMigrate(&Models.Email{})
	db.AutoMigrate(&Models.Adresse{})
	db.AutoMigrate(&Models.Collaborateur{})

	// Begin transation
	tx := db.Begin()

	// Clear current data
	tx.Unscoped().Delete(&Models.Depute{})
	tx.Unscoped().Delete(&Models.Site{})
	tx.Unscoped().Delete(&Models.Email{})
	tx.Unscoped().Delete(&Models.Adresse{})
	tx.Unscoped().Delete(&Models.Collaborateur{})

	// Inserting deputes
	for _, depute := range getDeputies() {
		depute.Depute.ID = 0
		tx.Create(&depute.Depute)
	}

	// Committing transaction
	tx.Commit()
}
