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
	//Get
	if r.Method == http.MethodGet {

	}

	url := "https://jsonplaceholder.typicode.com/todos"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("userId", "1")
	request.URL.RawQuery = params.Encode()

	fmt.Println(request.URL.String()) //https://jsonplaceholder.typicode.com/todos?userId=1

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
