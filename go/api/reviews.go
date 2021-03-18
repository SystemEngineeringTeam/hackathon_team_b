package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"with_b/db"
)

//Review は返す
func Review(w http.ResponseWriter, r *http.Request) {

	//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	params := r.URL.Query()

	if r.Method == http.MethodGet {
		var params_int int
		params_int, _ = strconv.Atoi(params["indexLectureNumber"][0])
		//reviewsは構造体のスライス
		reviews, err := db.CallReview(params_int)
		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
		//構造体からJSON文字列への変換する
		reviewsBytes, err := json.Marshal(reviews)
		fmt.Println("reviewsByte", reviewsBytes)
		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
		//stringに変換
		reviewsString := string(reviewsBytes)
		fmt.Fprintln(w, reviewsString)
	}

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
