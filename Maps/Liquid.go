package Maps

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/osteele/liquid"
)

func MapActivities(activities map[string]interface{}) string {
	engine := liquid.NewEngine()
	absPath, _ := filepath.Abs("./Maps/Activities.liquid")
	template, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(template))
	out, err := engine.ParseAndRenderString(string(template), activities)
	if err != nil {
		log.Fatalln(err)
	}

	return out
}
