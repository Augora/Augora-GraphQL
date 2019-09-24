package Importers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Augora/Augora-GraphQL/Maps"
	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func getDeputies() []Models.DeputeHandler {
	log.Println("Getting deputies...")
	deputesResp, err := http.Get("https://www.nosdeputes.fr/deputes/json")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Deputies received!")

	log.Println("Getting in office...")
	deputesEnMandatResp, err := http.Get("https://www.nosdeputes.fr/deputes/enmandat/json")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Deputies in office received!")

	var deputes Models.Deputes
	json.NewDecoder(deputesResp.Body).Decode(&deputes)

	var deputesEnMandat Models.Deputes
	json.NewDecoder(deputesEnMandatResp.Body).Decode(&deputesEnMandat)

	for deputeIndex := range deputes.Deputes {
		log.Println("Getting " + deputes.Deputes[deputeIndex].Depute.Slug + " activities...")
		activities := getDeputyActivities(deputes.Deputes[deputeIndex].Depute.Slug)
		deputes.Deputes[deputeIndex].Depute.Activites = activities
		log.Println(deputes.Deputes[deputeIndex].Depute.Slug + " received!")
		for _, deputeEnMandat := range deputesEnMandat.Deputes {
			if deputes.Deputes[deputeIndex].Depute.Slug == deputeEnMandat.Depute.Slug {
				deputes.Deputes[deputeIndex].Depute.EstEnMandat = true
			}
		}
	}

	return deputes.Deputes
}

func getDeputyActivities(slug string) []Models.Activity {
	activitesResp, err := http.Get("https://www.nosdeputes.fr/" + slug + "/graphes/lastyear/total?questions=true&format=json")
	if err != nil {
		log.Println(err)
	}

	bodyBytes, err := ioutil.ReadAll(activitesResp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var activitesFromAPI map[string]interface{}
	json.Unmarshal(bodyBytes, &activitesFromAPI)
	mappedActivities := Maps.MapActivities(activitesFromAPI)

	var activities []Models.Activity
	json.NewDecoder(strings.NewReader(mappedActivities)).Decode(&activities)

	return activities
}

func ImportDeputies() {
	db := Utils.GetDataBaseConnection()

	// Loading database models
	db.AutoMigrate(&Models.Depute{})
	db.AutoMigrate(&Models.Site{})
	db.AutoMigrate(&Models.Email{})
	db.AutoMigrate(&Models.Adresse{})
	db.AutoMigrate(&Models.Collaborateur{})
	db.AutoMigrate(&Models.Activity{})

	// Begin transation
	tx := db.Begin()

	// Clear current data
	tx.Unscoped().Delete(&Models.Depute{})
	tx.Unscoped().Delete(&Models.Site{})
	tx.Unscoped().Delete(&Models.Email{})
	tx.Unscoped().Delete(&Models.Adresse{})
	tx.Unscoped().Delete(&Models.Collaborateur{})
	tx.Unscoped().Delete(&Models.Activity{})

	// Inserting deputes
	for _, depute := range getDeputies() {
		depute.Depute.ID = 0
		tx.Create(&depute.Depute)
	}

	// Committing transaction
	tx.Commit()
}
