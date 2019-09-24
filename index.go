package main

import (
	"github.com/Augora/Augora-GraphQL/Handler"
	"github.com/Augora/Augora-GraphQL/Handler/Rest"
	"github.com/Augora/Augora-GraphQL/Importers"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("IMPORT") == "True" {
		Importers.ImportDeputies()
		Importers.ImportFiles()
	} else {
		http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
		http.HandleFunc("/deputes", Rest.Deputies)
		http.HandleFunc("/deputesenmandat", Rest.DeputiesInOffice)
		http.ListenAndServe(":3030", nil)
	}
}
