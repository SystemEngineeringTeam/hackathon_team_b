package api

import (
	"encoding/json"
	"fmt"
	"net/http"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
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
<<<<<<< HEAD
	// params.Add("grade", "3")
	// params.Add("department", "kk")
	// fmt.Println(params["grade"][0])
	// fmt.Println(params["department"][0])
	// fmt.Println(params)
=======
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

	//Get
	if r.Method == http.MethodGet {

<<<<<<< HEAD
		//recsは構造体のスライス
		recs, err := db.CallRectures(params["grade"][0], params["department"][0], params["semester"][0], params["dayofWeek"][0], params["time"][0], params["teacher"][0], params["lectureName"][0])
=======
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
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
<<<<<<< HEAD

		//byte型に変換する
		recBytes, err := json.Marshal(recs)

=======
		recBytes, err := json.Marshal(lectures)
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
		//エラー処理
		if err != nil {
			fmt.Println(err)
			return
		}
<<<<<<< HEAD

=======
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
		//stringに変換
		stringrecs := string(recBytes)
		fmt.Fprintln(w, stringrecs)
	}

<<<<<<< HEAD
	//lecture?grade="3"&department="kk"

}

func toJSON(r interface{}) string {
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ",")
=======
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
}
