package main

import (
	"net/http"
	"with_b/api"
	"with_b/db"
)

func main() {
	api.Test()

	// db.CallRectures()
	// db.CallReview(1)
	// db.CalculateStarAvarage(1)
	db.CallRectures("1","kk","前期","月曜","3限","高木","日本")
	http.HandleFunc("/lecture", api.Lectures)
	http.HandleFunc("/review",api.Review)

	http.ListenAndServe(":8080", nil)
}
