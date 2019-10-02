package Rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func DeputyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	db := Utils.GetDataBaseConnection()

	var depute Models.Depute
	slug := r.URL.Query()["slug"][0]
	queryResult := db.Preload("SitesWeb").Preload("Emails").Preload("Adresses").Preload("Collaborateurs").Where(&Models.Depute{Slug: slug}).Find(&depute)
	errors := queryResult.GetErrors()
	if len(errors) > 0 {
		fmt.Println("Some errors occured during the request:", errors)
		w.WriteHeader(http.StatusInternalServerError)
		var tmpErrors []struct {
			Error string `json:"error"`
		}
		for _, err := range errors {
			var tmp struct {
				Error string `json:"error"`
			}
			tmp.Error = err.Error()
			tmpErrors = append(tmpErrors, tmp)
		}
		res, _ := json.Marshal(tmpErrors)
		fmt.Fprintf(w, string(res))
		fmt.Println("Errors sent")
	} else {
		fmt.Println("Everything went well sending response...")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(depute)
		fmt.Fprintf(w, string(res))
		fmt.Println("Response sent")
	}
}
