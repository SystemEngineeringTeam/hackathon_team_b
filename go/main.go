package main

import (
	"net/http"
	"with_b/api"
	"with_b/db"
)

func main() {
	api.Test()
	db.CallRectures()

	http.HandleFunc("/lecture", api.Lectures)
	http.ListenAndServe(":8080", nil)
}
