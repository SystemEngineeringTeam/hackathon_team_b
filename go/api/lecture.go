package api

import (
	"fmt"
	"net/http"
	"with_b/db"
)

//Test はテストする関数
func Test(){
	fmt.Println("test")
}

func Lectures(w http.ResponseWriter,r *http.Request){

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")


	//Get
	if r.Method==http.MethodGet{

	}


}




