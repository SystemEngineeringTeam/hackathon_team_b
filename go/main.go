package main

import (
	"net/http"
	"with_b/api"
)

func main() {
	api.Test()

	// db.CallRectures()
	// db.CallReview(1)
	// db.CalculateStarAvarage(1)
	
	// fmt.Println(db.CallRectures("","kk","","","","",""))

	http.HandleFunc("/lecture", api.Lectures)
	http.HandleFunc("/review", api.Review)

	http.ListenAndServe(":3030", nil)
}
