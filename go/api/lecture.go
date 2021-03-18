package api

import (
	"encoding/json"
	"fmt"
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

	//クエリパラメータ
	params := r.URL.Query()

	//Get
	if r.Method == http.MethodGet {

		if params["grade"][0]=="null"{
			delete(params,"grade")
			params.Add("grade","")
		}
		if params["Department"][0]=="null"{
			delete(params,"Department")
			params.Add("Department","")
		}
		if params["semester"][0]=="null"{
			delete(params,"semester")
			params.Add("semester","")
		}
		if params["dayofweek"][0]=="null"{
			delete(params,"dayofweek")
			params.Add("dayofweek","")
		}
		if params["time"][0]=="null"{
			delete(params,"time")
			params.Add("time","")
		}
		if params["teacher"][0]=="null"{
			delete(params,"teacher")
			params.Add("teacher","")
		}
		if params["lectureName"][0]=="null"{
			delete(params,"lectureName")
			params.Add("lectureName","")
		}


		fmt.Println(params)

		lectures,err:=db.CallRectures(params["grade"][0],params["Department"][0],params["semester"][0],params["dayofweek"][0],params["time"][0],params["teacher"][0],params["lectureName"][0])

		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
		recBytes, err := json.Marshal(lectures)
		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
		//stringに変換
		stringrecs := string(recBytes)
		fmt.Fprintln(w, stringrecs)
	}

}
