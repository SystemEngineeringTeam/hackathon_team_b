package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Recture はjsonの構造体
type Recture struct{
	Grade string `json:"grade"`
	Department string `json:"department"`
	Semester string `json:"semester"`
	DayofWeek string `json:"dayofweek"`
	Time string `json:"time"`
	Teacher string `json:"teacher"`
	LectureName string `json:"lectureName"`
	ReviewStarAverage int `json:"reviewStarAverage"`
	IndexNumber int `json:"indexNumber"`

}

//grade string,department string,semester string,dayofWeek string,time string,teacher string,lectureName string
//CallRectures は講義の内容を呼び出す
func CallRectures()([]Recture,error){

	db, err := sql.Open("mysql", "root@/with_b")
    if err != nil {
        log.Println(err)
    }
	defer db.Close()

	//条件に合うものを取得する
	gr:="1"
	de:="kk"
	sem:="前期"
	day:="月曜"
	ti:="3限"
	te:="高木"
	le:="日本"

	//Queryを使えば複数のレコードを取得できる
	rows,err := db.Query("SELECT * FROM rectures WHERE grade=? AND department=? AND semester=? AND dayofweek=? AND time=? AND teacher like '%'||?||'%' AND lectureName like '%'||?||'%'",gr,de,sem,day,ti,te,le)
    defer rows.Close()
    if err != nil {
        log.Println(err)
    }

	//空の構造体のスライスを作成
	Rectures:=make([]Recture,0,0)

	for rows.Next(){
		var grade string
		var department string
		var semester string
		var dayofweek string
		var time string
		var teacher string
		var lectureName string
		var reviewStarAverage int
		var indexNumber int

		if err:=rows.Scan(&grade,&department,&semester,&dayofweek,&time,&teacher,&lectureName,&reviewStarAverage,&indexNumber);err!=nil{
			log.Print(err)
		}

		//構造体に格納
		tmpRecture:=Recture{grade,department,semester,dayofweek,time,teacher,lectureName,reviewStarAverage,indexNumber}

		//スライスに追加
		Rectures=append(Rectures,tmpRecture)
	}

	fmt.Println(Rectures)

	return Rectures,nil
}
