package api

import "net/http"

//Review は返す
func Review(w http.ResponseWriter,r *http.Request){

		//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")


	if r.Method==http.MethodGet{

	}


	if r.Method==http.MethodPost{

	}



}