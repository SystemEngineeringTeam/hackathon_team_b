package api

import (
	"fmt"
	"log"
	"net/http"
	"with_b/db"
)

//Test はテストする関数
func Test() {
	fmt.Println("test")
}

func Lectures(w http.ResponseWriter, r *http.Request) {

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	db.CallRectures(" ", " ", " ", " ", " ", " ", " ")
	url := "https://jsonplaceholder.typicode.com/todos"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//lecture?grade="3"&department="kk"
	//クエリパラメータ
	params := r.URL.Query()
	params.Add("grade", "3")
	params.Add("department", "kk")
	fmt.Println(params["grade"][0])
	fmt.Println(params["department"][0])

}

//Get
// if r.Method == http.MethodGet {

// }
