package main

import (
	"fmt"
	"net/http"
	"with_b/api"
	"with_b/db"
)

func main() {
	api.Test()

	// db.CallRectures()
	// db.CallReview(1)
	// db.CalculateStarAvarage(1)

	fmt.Println(db.CallRectures("1","kk","","","","","日本"))
	// fmt.Println(db.CallRectures("3","kk","後期","火曜","4限","中村","情報"))

	http.HandleFunc("/lecture", api.Lectures)
	http.HandleFunc("/review",api.Review)

	http.ListenAndServe(":8080", nil)
}
