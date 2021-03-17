package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"with_b/db"
)

//Review は返す
func Review(w http.ResponseWriter, r *http.Request) {

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// if r.Method==http.MethodGet{

	// }

	// var body = []byte(`[
	//     {"Name": "Platypus", "Order": "Monotremata",sentence:"soufasfaof"},
	// ]`)

	//8080/review
	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//ReviewDetailの構造体の初期化
		jsonreview := db.ReviewDetail{}
		//構造体にbody中身を入れる
		//json.Unmarshalして構造体に変換する

		err = json.Unmarshal(body, &jsonreview)
		if err != nil {
			fmt.Println("error:", err)
		}

		//登録処理
		err = db.RegisterReview(jsonreview)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Fprintln(w, jsonreview)
	}

}
