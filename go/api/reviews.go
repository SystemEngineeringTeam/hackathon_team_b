package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"with_b/db"
)

//Review は返す
func Review(w http.ResponseWriter,r *http.Request){

		//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	params := r.URL.Query()

    if r.Method == http.MethodGet {
        var paramsInt int
        paramsInt,_= strconv.Atoi(params["indexLectureNumber"][0])

        //reviewsは構造体のスライス
        reviews, err := db.CallReview(paramsInt)
        //エラー処理
        if err != nil {
            fmt.Println(err)
            return
        }
        //構造体からJSON文字列への変換する
        reviewsBytes, err := json.Marshal(reviews)
        //エラー処理
        if err != nil {
            fmt.Println(err)
            return
        }
        //stringに変換
        reviewsString := string(reviewsBytes)
        fmt.Fprintln(w, reviewsString)
    }

	if r.Method==http.MethodPost{
		//body読み込み
		jsonBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("io error")
			return
		}

		//構造体の初期化
		data:= db.ReviewDetail{}

		//taskの構造体にbodyの値を入れる
		if err := json.Unmarshal(jsonBytes,&data); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Println("JSON Unmarshal error:",err)
			return
		}
		fmt.Println(data)

		err=db.RegisterReview(data)

		if err!=nil{
			log.Println(err)
		}
		fmt.Fprintln(w,"hello")
	}
}