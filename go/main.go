package main

import (
	"net/http"
	"with_b/api"
	"with_b/db"
)



func main()  {
	api.Test()


	// db.CallRectures()
	//db.CallReview(1)
	// db.CalculateStarAvarage(1)
	db.RegisterReview(1)

	http.HandleFunc("/lecture",api.Lectures)
	http.ListenAndServe(":8080",nil)
}