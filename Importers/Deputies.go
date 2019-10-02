package Importers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
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

func setField(v interface{}, name string, value string) error {
	// v must be a pointer to a struct
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	// Field must be exported
	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	// We expect a string field
	if fv.Kind() != reflect.String {
		return fmt.Errorf("%s is not a string field", name)
	}

	// Set the value
	fv.SetString(value)
	return nil
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
	defer db.Close()

	// Loading database models
	// db.AutoMigrate(&Models.Depute{})
	// db.AutoMigrate(&Models.Site{})
	// db.AutoMigrate(&Models.Email{})
	// db.AutoMigrate(&Models.Adresse{})
	// db.AutoMigrate(&Models.Collaborateur{})
	// db.AutoMigrate(&Models.Activity{})

	// Begin transation
	// tx := db.Begin()

	var deputiesInDB []Models.Depute
	db.Set("gorm:auto_preload", true).Find(&deputiesInDB)

	deputies := getDeputies()
	deputies[0].GroupeSigle = "LEL"
	deputies = append(deputies, Models.Depute{Slug: "mdr"})

	// Inserting deputes
	changelog, _ := diff.Diff(deputiesInDB, deputies)
	for _, change := range changelog {
		if change.Type == "update" {
			deputyIndex, _ := strconv.Atoi(change.Path[0])
			currentDeputy := deputiesInDB[deputyIndex]
			setField(currentDeputy, change.Path[1], change.To.(string))
		}
	}
	jsonContent, _ := json.MarshalIndent(changelog, "", "  ")
	jsonString := string(jsonContent)
	fmt.Println(jsonString)
	db.Save(&deputiesInDB)

	// Committing transaction
	// tx.Commit()
}
