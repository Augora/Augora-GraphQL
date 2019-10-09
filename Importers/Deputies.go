package Importers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Augora/Augora-GraphQL/Maps"
	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
	"github.com/r3labs/diff"
)

func getDeputies() []Models.Depute {
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

	database := os.Getenv("BACKEND_SQL_DATABASE")
	if database == "sandbox-db" {
		deputes.Deputes = deputes.Deputes[:20]
	}

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

	var res []Models.Depute
	for _, depute := range deputes.Deputes {
		res = append(res, depute.Depute)
	}

	return res
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

func DiffFromDB(fromDB []Models.Depute, fromAPI []Models.Depute) []Models.DeputyDiff {
	var res []Models.DeputyDiff

	changelog, _ := diff.Diff(fromDB, fromAPI)
	groupedDiff := make(map[string]diff.Changelog)
	for _, change := range changelog {
		groupedDiff[change.Path[0]] = append(groupedDiff[change.Path[0]], change)
	}

	for slug := range groupedDiff {
		// changedGroup := groupedDiff[slug]
		var deputyInDB Models.Depute
		foundDeputyInDB := false
		for i := range fromDB {
			if fromDB[i].Slug == slug {
				deputyInDB = fromDB[i]
				foundDeputyInDB = true
				break
			}
		}
		// Check if deputy was found
		if foundDeputyInDB {
			foundDeputyInAPI := false
			var deputyInAPI Models.Depute
			for i := range fromAPI {
				if fromAPI[i].Slug == slug {
					deputyInAPI = fromAPI[i]
					foundDeputyInAPI = true
					break
				}
			}
			if foundDeputyInAPI {
				// Update
				updatedDeputy := Models.MergeDeputies(deputyInDB, deputyInAPI)
				newDiff := Models.DeputyDiff{
					Operation: "Update",
					Deputy:    updatedDeputy,
				}
				res = append(res, newDiff)
			} else {
				// Delete
			}
		} else {
			// Create
			var newDeputy Models.Depute
			for i := range fromAPI {
				if fromAPI[i].Slug == slug {
					newDeputy = fromAPI[i]
					break
				}
			}
			newDiff := Models.DeputyDiff{
				Operation: "Create",
				Deputy:    newDeputy,
			}
			res = append(res, newDiff)
		}
	}

	return res
}

func ImportDeputies() {
	db := Utils.GetDataBaseConnection()
	defer db.Close()

	// Begin transation
	tx := db.Begin()

	// Loading database models
	tx.AutoMigrate(&Models.Depute{})
	tx.AutoMigrate(&Models.Site{})
	tx.AutoMigrate(&Models.Email{})
	tx.AutoMigrate(&Models.Adresse{})
	tx.AutoMigrate(&Models.Collaborateur{})
	tx.AutoMigrate(&Models.Activity{})

	deputies := getDeputies()
	deputies[0].GroupeSigle = "FI"
	var deputiesInDB []Models.Depute
	db.Set("gorm:auto_preload", true).Find(&deputiesInDB)

	diffs := DiffFromDB(deputiesInDB, deputies)

	for _, diff := range diffs {
		if diff.Operation == "Create" {
			tx.Create(diff.Deputy)
		}
		if diff.Operation == "Update" {
			tx.Save(diff.Deputy)
		}
	}

	// Committing transaction
	// tx.Commit()
}
