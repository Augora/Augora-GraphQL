package main

import (
	"github.com/Augora/Augora-GraphQL/Importers"
	"os"
	"net/http"
	"github.com/Augora/Augora-GraphQL/Handler"
)

func main() {
	if os.Getenv("IMPORT") == "True" {
		Importers.ImportDeputies()
		Importers.ImportFiles()
	} else {
		http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
		http.ListenAndServe(":3030", nil)
	}
}
