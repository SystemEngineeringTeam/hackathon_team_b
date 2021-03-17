package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ReviewDetail struct{
	ReviewStar int `json:"department"`
	Contents string `json:"contents"`
}



//CallReview はレビューの構造体のスライスを返す
func CallReview(indexLectureNumber int)([]ReviewDetail,error)  {


	db, err := sql.Open("mysql", "root@/with_b")
    if err != nil {
        log.Println(err)
		return nil,err
    }
	defer db.Close()


	rows,err:=db.Query("SELECT * FROM reviews WHERE indexLectureNumber=?",indexLectureNumber)
	defer rows.Close()
    if err != nil {
        log.Println(err)
		return nil,err
    }

	ReviewDetails:=make([]ReviewDetail,0,0)


	for rows.Next(){

		var indexLectureNumber int
		var star int
		var contents string
		var id int

		if err:=rows.Scan(&indexLectureNumber,&star,&contents,&id);err!=nil{
			log.Print(err)
			return nil,err
		}

		tmpReview:=ReviewDetail{star,contents}

		ReviewDetails=append(ReviewDetails,tmpReview)
	}

	fmt.Println(ReviewDetails)

	return ReviewDetails,nil
}

//CalculateStarAvarage は星の平均を数える
func CalculateStarAvarage(indexLectureNumber int)(string,error){
	db, err := sql.Open("mysql", "root@/with_b")
    if err != nil {
        log.Println(err)
		return "",err
    }
	defer db.Close()

	rows,err:=db.Query("SELECT reviewStar FROM reviews WHERE indexLectureNumber=?",indexLectureNumber)
    if err != nil {
        log.Println(err)
		return "",err
    }
	defer rows.Close()

	var starTotal float64
	starTotal=0

	rowsCount:=0

	for rows.Next(){
		var star float64

		if err:=rows.Scan(&star)
		err!=nil{
			log.Print(err)
			return "",err
		}
		starTotal=starTotal+star
		rowsCount++
	}


	starAvarage:=starTotal/float64(rowsCount)

	fmt.Println(starTotal)
	fmt.Println(starAvarage)

	stringStarAvarage:=strconv.FormatFloat(starAvarage, 'f', 2, 64)

	fmt.Println(stringStarAvarage)

	return stringStarAvarage,nil
}


//RegisterReview はレビューを登録する
func RegisterReview(indexLectureNumber int){

	db, err := sql.Open("mysql", "root@/with_b")
    if err != nil {
        log.Println(err)
    }
	defer db.Close()

	_,err=db.Exec("INSERT INTO reviews (indexLectureNumber,reviewStar,sentence) VALUES ('1','5','hogehogegehoge')")

	if err != nil {
        log.Println(err)
    }

}






