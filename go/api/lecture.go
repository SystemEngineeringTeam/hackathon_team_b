package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	url := "https://jsonplaceholder.typicode.com/todos"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//クエリパラメータ
	params := r.URL.Query()
	params.Add("grade", "3")
	params.Add("department", "kk")
	fmt.Println(params["grade"][0])
	fmt.Println(params["department"][0])
	// fmt.Println(params)

	//Get
	if r.Method == http.MethodGet {

		//recsは構造体のスライス
		recs, err := db.CallRectures(params["grade"][0], params["department"][0], params["semester"][0], params["dayofWeek"][0], params["time"][0], params["teacher"][0], params["lectureName"][0])

		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}

		//byte型に変換する
		recBytes, err := json.Marshal(recs)

		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}

		//stringに変換
		stringrecs := string(recBytes)
		fmt.Fprintln(w, stringrecs)
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
