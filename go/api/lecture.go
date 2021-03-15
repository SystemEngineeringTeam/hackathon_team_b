package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	//クエリパラメータ
	params := r.URL.Query()
	params.Add("userId", "1")

	fmt.Println(r.URL.String())
	}

	//Get
	if r.Method == http.MethodGet {

	}

