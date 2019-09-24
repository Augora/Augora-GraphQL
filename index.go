package main

import (
	"net/http"
	"os"

	"github.com/Augora/Augora-GraphQL/Handler"
	"github.com/Augora/Augora-GraphQL/Handler/Rest"
	"github.com/Augora/Augora-GraphQL/Importers"
)

func main() {
	if os.Getenv("IMPORT") == "True" {
		Importers.ImportDeputies()
		Importers.ImportFiles()
	} else {
		http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
		http.HandleFunc("/deputes", Rest.DeputiesHandler)
		http.HandleFunc("/deputesenmandat", Rest.DeputiesInOfficeHandler)
		http.ListenAndServe(":3030", nil)
	}
}
