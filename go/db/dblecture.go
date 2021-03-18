package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Recture はjsonの構造体
type Recture struct {
	Grade              string `json:"grade"`
	Department         string `json:"department"`
	Semester           string `json:"semester"`
	DayofWeek          string `json:"dayofweek"`
	Timed              string `json:"timed"`
	Teacher            string `json:"teacher"`
	LectureName        string `json:"lectureName"`
	ReviewStarAverage  string `json:"reviewStarAverage"`
	IndexLectureNumber int    `json:"indexLectureNumber"`
}

//gr string, de string, sem string, day string, ti string, te string, le string
//CallRectures は講義の内容を呼び出す
func CallRectures(gr string, de string, sem string, day string, ti string, te string, le string) ([]Recture, error) {

	db, err := sql.Open("mysql", "root:tako64tako@tcp(127.0.0.1:3306)/with_b")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	//条件に合うものを取得する
	// gr:="1"
	// de:="kk"
	// sem:="前期"
	// day:="月曜"
	// ti:="3限"
	// te:="高木"
	// le:="日本"
	// if gradeFlag{

	// }

	//クエリの有無

	//indexLectureNumberは取ってこれる
	//Queryを使えば複数のレコードを取得できる

	sql := "SELECT * FROM lectures WHERE "
	if len(de) != 0 {
		sql = sql + "department=" + "'" + de + "'" + " AND "
	}
	if len(sem) != 0 {
		sql = sql + "semester=" + "'" + sem + "'" + " AND "
	}
	if len(day) != 0 {
		sql = sql + "dayofweek=" + "'" + day + "'" + " AND "
	}
	if len(ti) != 0 {
		sql = sql + "timed=" + "'" + ti + "'" + " AND "
	}
	if len(te) != 0 {
		sql = sql + "teacher like " + "'%" + te + "%'" + " AND "
	}
	if len(le) != 0 {
		sql = sql + "lectureName like " + "'%" + le + "%'"
	}

	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	//空の構造体のスライスを作成
	Rectures := make([]Recture, 0, 0)

	for rows.Next() {
		// var grade string
		var department string
		var semester string
		var dayofweek string
		var timed string
		var teacher string
		var lectureName string
		var indexLectureNumber int

		rows.Scan(&department, &semester, &dayofweek, &timed, &teacher, &lectureName, &indexLectureNumber)

		reviewStarAverage, err := CalculateStarAvarage(indexLectureNumber)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("indexLecturenumber", indexLectureNumber)

		gradeFlag := false
		row, err := db.Query("SELECT grade FROM grade WHERE indexLectureNumber=?", indexLectureNumber)
		defer row.Close()

		var gradeSlice string

		for row.Next() {
			var tmpGrade string
			row.Scan(&tmpGrade)
			if tmpGrade == gr {
				gradeFlag = true
			}
			gradeSlice = gradeSlice + tmpGrade + " "
		}

		//一致した学年が一つでもあればindexLectureNumberに対応する学年全て返す
		fmt.Println(gradeSlice)

		if gradeFlag {
			//構造体に格納
			tmpRecture := Recture{gradeSlice, department, semester, dayofweek, timed, teacher, lectureName, reviewStarAverage, indexLectureNumber}
			//スライスに追加
			Rectures = append(Rectures, tmpRecture)
		} else {
			return nil, err
		}

	}

	fmt.Println(Rectures)

	return Rectures, nil
}

//関数化
