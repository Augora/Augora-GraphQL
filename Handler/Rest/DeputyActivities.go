package Rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func ActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	db := Utils.GetDataBaseConnection()

	var depute Models.Depute
	slug := r.URL.Query()["slug"][0]
	queryResult := db.
		Preload("Activites").
		Where(&Models.Depute{Slug: slug}).
		Find(&depute)
	errors := queryResult.GetErrors()
	var count int
	queryResult.Count(&count)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(errors)
		fmt.Fprintf(w, string(res))
	} else {
		if count > 0 {
			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal(depute.Activites)
			fmt.Fprintf(w, string(res))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			var tmp struct {
				error string `json:"error"`
			}
			tmp.error = "This deputy does not exists."
			res, _ := json.Marshal(tmp)
			fmt.Fprintf(w, string(res))
		}
	}
}
