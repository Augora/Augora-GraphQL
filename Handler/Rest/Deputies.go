package Rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Models"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func DeputiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "s-maxage=3600")
	db := Utils.GetDataBaseConnection()

	var deputes []Models.Depute
	db.Preload("sites").Preload("emails").Preload("adresses").Preload("collaborateurs").Find(&deputes)
	res, _ := json.Marshal(deputes)
	fmt.Fprintf(w, string(res))
}
