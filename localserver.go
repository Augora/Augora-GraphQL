package main

import (
  "net/http"
  
	"github.com/KevinBacas/Gin-Go-Test/Handler"
)

func main() {
  http.HandleFunc("/graphql", handler.GraphQLHTTPHandler)
  http.ListenAndServe(":3030", nil)
}