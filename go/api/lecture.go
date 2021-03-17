package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	//クエリパラメータ
	params := r.URL.Query()
	params.Add("grade", "3")
	params.Add("department", "kk")
	fmt.Println(params["grade"][0])
	fmt.Println(params["department"][0])
	fmt.Println(params)

	//Get
	if r.Method == http.MethodGet {


		// エラー処理
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// byte型に変換する
		// recBytes, err := json.Marshal(recs)

		// エラー処理
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// stringに変換
		// stringrecs := string(recBytes)
		// fmt.Fprintln(w, stringrecs)
	}

	url := "https://jsonplaceholder.typicode.com/todos"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//lecture?grade="3"&department="kk"

}

func toJSON(r interface{}) string {
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ",")
}
