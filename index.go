package main

import (
	"net/http"

	Handler "github.com/Augora/Augora-GraphQL/Handler"
)

func main() {
	http.HandleFunc("/graphql", Handler.GraphQLHTTPHandler)
	http.ListenAndServe(":3030", nil)
}
