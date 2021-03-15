package main

import (
	"net/http"
	"with_b/api"
)

func main() {
	api.Test()

	http.HandleFunc("/lecture", api.Lectures)
	http.ListenAndServe(":8080", nil)
}
