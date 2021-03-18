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

	// params := r.URL.Query()


	//Get
	if r.Method == http.MethodGet {


		// params["department"][0]=""
		// params["semester"][0]=""
		// params["dayofweek"][0]=""
		// params["time"][0]=""
		// params["teacher"][0]=""
		// params["lectureName"][0]=""

		// _,ok:=params["grade"]
		// if ok{
		// 	params["grade"][0]=""
		// 	fmt.Println("okg")
		// }else{
		// 	fmt.Println("nog")
		// }

		// if _,ok:=params["semester"];!ok{
		// 	params["semester"][0]=""
		// 	fmt.Println("No se")
		// }
		// if _,ok:=params["time"];!ok{
		// 	params["time"][0]=""
		// 	fmt.Println("No ti")
		// }
		// if _,ok:=params["teacher"];!ok{
		// 	params["teacher"][0]=""
		// 	fmt.Println("No te")
		// }
		// if _,ok:=params["dayofweek"];!ok{
		// 	params["dayofweek"][0]=""
		// 	fmt.Println("No dayof")
		// }
		// if _,ok:=params["lectureName"];!ok{
		// 	params["lectureName"][0]=""
		// 	fmt.Println("No leNa")
		// }
		// if _,ok:=params["department"];!ok{
		// 	params["department"][0]=""
		// 	fmt.Println("No depa")
		// }

		fmt.Println("hello")
		// fmt.Println(params["grade"][0],params["department"][0],params["semester"][0],params["dayofweek"][0],params["time"][0],params["teacher"][0],params["lectureName"][0])

		// lectures,err:=db.CallRectures(params["grade"][0],params["department"][0],params["semester"][0],params["dayofweek"][0],params["time"][0],params["teacher"][0],params["lectureName"][0])

		fmt.Println("bye")
		lectures,err:=db.CallRectures("1","kk","前期","月曜","３限","高木健太郎","日本")

		// http://localhost:8080/lecture?grade=1&"department=kk&semester=前期&dayofweek=月曜&time=３限&teacher=高木&lectureName=日本
		// fmt.Println(params["grade"][0])
		// fmt.Println(params["department"][0])
		// fmt.Println(params["semester"][0])
		// fmt.Println(params["dayofweek"][0])
		// fmt.Println(params["time"][0])
		// fmt.Println(params["teacher"][0])
		// fmt.Println(params["lectureName"][0])

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

func toJSON(r interface{}) string {
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ",")
}
