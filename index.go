package main

import (
	"net/http"

	handler "github.com/Augora/Augora-GraphQL/Handler"
)

func main() {
	http.HandleFunc("/graphql", handler.GraphQLHTTPHandler)
	http.ListenAndServe(":3030", nil)
}
