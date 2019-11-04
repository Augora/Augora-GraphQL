package Maps

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"encoding/json"
	"strings"

	"github.com/osteele/liquid"
	"github.com/Augora/Augora-GraphQL/Models"
)

func MapActivities(activities map[string]interface{}) Models.ActivitesHandler {
	// Loading liquid map
	engine := liquid.NewEngine()
	absPath, _ := filepath.Abs("./Maps/Activities.liquid")
	template, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Passing raw activites through the liquid map
	mappedActivities, err := engine.ParseAndRenderString(string(template), activities)
	if err != nil {
		log.Fatalln(err)
	}

	// Decode activites string to Go Object
	var activitiesHandler Models.ActivitesHandler
	err = json.NewDecoder(strings.NewReader(mappedActivities)).Decode(&activities)
	if err != nil {
		log.Fatalln(err)
	}

	return activitiesHandler
}
