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
	// params := r.URL.Query()

	//Get
	if r.Method == http.MethodGet {

		// lectures,err:=db.CallRectures(params["grade"][0],params["department"][0],params["semester"][0],params["dayofweek"][0],params["time"][0],params["teacher"][0],params["lectureName"][0])

		lectures, err := db.CallRectures("", "", "", "", "", "", "")

		// http://localhost:8080/lecture?grade=1&"department=kk&semester=前期&dayofweek=月曜&time=３限&teacher=高木&lectureName=日本

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
