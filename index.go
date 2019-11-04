package main

import (
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Handler"
	"github.com/Augora/Augora-GraphQL/Handler/Rest"
	"github.com/Augora/Augora-GraphQL/Importers"
	"github.com/Augora/Augora-GraphQL/Utils"
)

func main() {
	if os.Getenv("IMPORT") == "True" {
		db := Utils.GetDataBaseConnection()
		defer db.Close()

		Importers.ImportDeputies(db)
		// Importers.ImportFiles()
	} else {
		http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
		http.HandleFunc("/deputes", Rest.DeputiesHandler)
		http.HandleFunc("/deputesenmandat", Rest.DeputiesInOfficeHandler)
		http.ListenAndServe(":3030", nil)
	}
}
