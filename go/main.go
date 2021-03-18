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
<<<<<<< HEAD
	// db.RegisterReview(1)
=======
	
	// fmt.Println(db.CallRectures("","kk","","","","",""))
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

	http.HandleFunc("/lecture", api.Lectures)
	http.HandleFunc("/review", api.Review)

<<<<<<< HEAD
	http.ListenAndServe(":8080", nil)
=======
	http.ListenAndServe(":3030", nil)
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
}
