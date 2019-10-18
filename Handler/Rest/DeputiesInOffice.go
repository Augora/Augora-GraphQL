package Rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func DeputiesInOfficeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	db := Utils.GetDataBaseConnection()

	var deputes []Models.Depute
	db.Preload("SitesWeb").Preload("Emails").Preload("Adresses").Preload("Collaborateurs").Preload("AnciensMandats").Preload("AutresMandats").Where(&Models.Depute{EstEnMandat: true}).Find(&deputes)
	res, _ := json.Marshal(deputes)
	fmt.Fprintf(w, string(res))
}
