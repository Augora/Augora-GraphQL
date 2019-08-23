package main

import (
	"net/http"

	"github.com/Augora/Augora-GraphQL/Handler"
)

func main() {
	// Importers.ImportDeputies()
	// Importers.ImportFiles()

	http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
	http.ListenAndServe(":3030", nil)
}
