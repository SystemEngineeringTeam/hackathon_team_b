package db

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

<<<<<<< HEAD
type ReviewDetail struct {
	IndexLectureNumber int    `json:"indexLectureNumber"`
	ReviewStar         int    `json:"reviewStar"`
	Contents           string `json:"sentence"`
}

=======
type ReviewDetail struct{
	IndexLectureNumber int `json:"indexLectureNumber"`
	ReviewStar int `json:"reviewStar"`
	Contents string `json:"sentence"`
}


>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
//CallReview はレビューの構造体のスライスを返す
func CallReview(indexLectureNumber int) ([]ReviewDetail, error) {

	db, err := sql.Open("mysql", "root:tako64tako@tcp(127.0.0.1:3306)/with_b")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

<<<<<<< HEAD
	rows, err := db.Query("SELECT * FROM reviews WHERE indexLectureNumber=?", indexLectureNumber)
=======
	rows,err:=db.Query("SELECT * FROM reviews WHERE indexLectureNumber=?",indexLectureNumber)
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

<<<<<<< HEAD
	ReviewDetails := make([]ReviewDetail, 0, 0)

	for rows.Next() {
=======
	for rows.Next(){
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

		var indexLectureNumber int
		var star int
		var sentence string
		var id int

<<<<<<< HEAD
		if err := rows.Scan(&indexLectureNumber, &star, &sentence, &id); err != nil {
=======
		if err:=rows.Scan(&indexLectureNumber,&star,&sentence,&id);err!=nil{
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400
			log.Print(err)
			return nil, err
		}

<<<<<<< HEAD
		tmpReview := ReviewDetail{indexLectureNumber, star, sentence}
=======
		tmpReview:=ReviewDetail{indexLectureNumber,star,sentence}
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

		ReviewDetails = append(ReviewDetails, tmpReview)
	}
	// fmt.Println(ReviewDetails)

	return ReviewDetails, nil
}

//CalculateStarAvarage は星の平均を数える
func CalculateStarAvarage(indexLectureNumber int) (string, error) {
	db, err := sql.Open("mysql", "root@/with_b")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer db.Close()

	rows, err := db.Query("SELECT reviewStar FROM reviews WHERE indexLectureNumber=?", indexLectureNumber)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rows.Close()

	var starTotal float64
	starTotal = 0

	rowsCount := 0

	for rows.Next() {
		var star float64

		if err := rows.Scan(&star); err != nil {
			log.Print(err)
			return "", err
		}
		starTotal = starTotal + star
		rowsCount++
	}

	starAvarage := starTotal / float64(rowsCount)

	// fmt.Println(starTotal)
	// fmt.Println(starAvarage)

	stringStarAvarage := strconv.FormatFloat(starAvarage, 'f', 2, 64)

	// fmt.Println(stringStarAvarage)

	return stringStarAvarage, nil
}

//RegisterReview はレビューを登録する
<<<<<<< HEAD
func RegisterReview(re ReviewDetail) error {
=======
func RegisterReview(re ReviewDetail)(error){
>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

	db, err := sql.Open("mysql", "root@/with_b")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	fmt.Println(re.ReviewStar)

<<<<<<< HEAD
	_, err = db.Exec("INSERT INTO reviews (indexLectureNumber,reviewStar,sentence) VALUES (?,?,?)", re.IndexLectureNumber, re.ReviewStar, re.Contents)
=======
	_,err=db.Exec("INSERT INTO reviews (IndexLectureNumber,ReviewStar,sentence) VALUES (?,?,?)",re.IndexLectureNumber,re.ReviewStar,re.Contents)


>>>>>>> a79a334143f1527c233a98f50eb5221bbee92400

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
