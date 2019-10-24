package Importers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

	log.Println("End getDeputies!")
	return res
}

func getDeputyActivities(slug string) []Models.Activite {
	activitesResp, err := http.Get("https://www.nosdeputes.fr/" + slug + "/graphes/lastyear/total?questions=true&format=json")
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err := ioutil.ReadAll(activitesResp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var activitesFromAPI map[string]interface{}
	json.Unmarshal(bodyBytes, &activitesFromAPI)
	mappedActivities := Maps.MapActivities(activitesFromAPI)

	var activities Models.ActivitesHandler
	err = json.NewDecoder(strings.NewReader(mappedActivities)).Decode(&activities)
	if err != nil {
		log.Fatalln(err)
	}

	layoutISO := "2006-01-02"
	t, err := time.Parse(layoutISO, activities.DateFin)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		if t.Weekday() == 1 {
			break
		}
		t = t.AddDate(0, 0, -1)
	}

	for i := range activities.Data {
		newStartDate := t.AddDate(0, 0, (int)(-(54-activities.Data[i].NumeroDeSemaine)*7))
		newEndDate := newStartDate.AddDate(0, 0, 7)
		activities.Data[i].DateDebut = newStartDate
		activities.Data[i].DateFin = newEndDate
	}

	return activities.Data
}

func DiffFromDB(fromDB []Models.Depute, fromAPI []Models.Depute) []Models.GenericDiff {
	var res []Models.GenericDiff

	changelog, _ := diff.Diff(fromDB, fromAPI)
	groupedDiff := make(map[string]diff.Changelog)
	for _, change := range changelog {
		groupedDiff[change.Path[0]] = append(groupedDiff[change.Path[0]], change)
	}

	jsonContent, _ := json.MarshalIndent(groupedDiff, "", "  ")
	jsonString := string(jsonContent)
	fmt.Println(jsonString)

	for slug := range groupedDiff {
		changedGroup := groupedDiff[slug]
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
				// Delete fields
				path := []string{slug, "Sites"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var siteToDelete Models.Site
						for _, s := range deputyInDB.Sites {
							if s.Site == change.Path[2] {
								siteToDelete = s
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      siteToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "Emails"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var emailToDelete Models.Email
						for _, e := range deputyInDB.Emails {
							if e.Email == change.Path[2] {
								emailToDelete = e
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      emailToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "Adresses"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var adresseToDelete Models.Adresse
						for _, a := range deputyInDB.Adresses {
							if a.Adresse == change.Path[2] {
								adresseToDelete = a
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      adresseToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "Collaborateurs"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var collabToDelete Models.Collaborateur
						for _, c := range deputyInDB.Collaborateurs {
							if c.Collaborateur == change.Path[2] {
								collabToDelete = c
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      collabToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "AnciensMandats"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var ancienMandatToDelete Models.AncienMandat
						for _, am := range deputyInDB.AnciensMandats {
							if am.AncienMandat == change.Path[2] {
								ancienMandatToDelete = am
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      ancienMandatToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "AutresMandats"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var autreMandatToDelete Models.AutreMandat
						for _, am := range deputyInDB.AutresMandats {
							if am.AutreMandat == change.Path[2] {
								autreMandatToDelete = am
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      autreMandatToDelete,
						}
						res = append(res, newDiff)
					}
				}

				path = []string{slug, "Activites"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						var activiteToDelete Models.Activite
						for _, a := range deputyInDB.Activites {
							numSemaine, _ := strconv.ParseUint(change.Path[2], 10, 32)
							var numSemaine32 uint
							numSemaine32 = uint(numSemaine)
							if a.NumeroDeSemaine == numSemaine32 {
								activiteToDelete = a
							}
						}
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      activiteToDelete,
						}
						res = append(res, newDiff)
					}
				}

				// Update
				updatedDeputy := Models.MergeDeputies(deputyInDB, deputyInAPI)
				newDiff := Models.GenericDiff{
					Operation: "Update",
					Item:      updatedDeputy,
				}
				res = append(res, newDiff)
			} else {
				// Delete Deputy
				path := []string{slug, "Slug"}
				for _, change := range changedGroup.Filter(path) {
					if change.Type == "delete" {
						newDiff := Models.GenericDiff{
							Operation: "Delete",
							Item:      deputyInDB,
						}
						res = append(res, newDiff)
					}
				}
			}
		} else {
			// Create
			var newDeputy Models.Depute
			for i := range fromAPI {
				if fromAPI[i].Slug == slug {
					newDeputy = fromAPI[i]
					newDeputy.ID = 0
					break
				}
			}
			newDiff := Models.GenericDiff{
				Operation: "Create",
				Item:      newDeputy,
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
	tx.AutoMigrate(&Models.AutreMandat{})
	tx.AutoMigrate(&Models.AncienMandat{})
	tx.AutoMigrate(&Models.Activite{})

	deputies := getDeputies()
	var deputiesInDB []Models.Depute
	errors := tx.Set("gorm:auto_preload", true).Find(&deputiesInDB).GetErrors()
	for _, err := range errors {
		fmt.Println(err)
	}
	diffs := DiffFromDB(deputiesInDB, deputies)
	jsonContent, _ := json.MarshalIndent(diffs, "", "  ")
	jsonString := string(jsonContent)
	fmt.Println(jsonString)

	for _, diff := range diffs {
		if diff.Operation == "Create" {
			tx.Create(diff.Item)
		}
		if diff.Operation == "Update" {
			tx.Save(diff.Item)
		}
		if diff.Operation == "Delete" {
			tx.Delete(diff.Item)
		}
	}

	// Committing transaction
	tx.Commit()
}
