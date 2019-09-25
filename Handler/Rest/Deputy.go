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
	db.Preload("SitesWeb").Preload("Emails").Preload("Adresses").Preload("Collaborateurs").Where(&Models.Depute{Slug: slug}).Find(&depute)
	res, _ := json.Marshal(depute)
	fmt.Fprintf(w, string(res))
}
