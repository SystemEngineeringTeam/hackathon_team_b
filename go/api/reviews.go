package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"with_b/db"
)

//Review は返す
func Review(w http.ResponseWriter,r *http.Request){

		//セキリティ設定
	w.Header().Set("Access-Control-Allow-Origin", "*")                       // Allow any access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method==http.MethodGet{
		fmt.Fprintln(w,"hello")
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