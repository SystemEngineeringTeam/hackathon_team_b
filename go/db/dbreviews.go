package db

import (
	_ "github.com/go-sql-driver/mysql"
)

//Recture はjsonの構造体
type Recture struct {
	Grade       string `json:"grade"`
	Department  string `json:"department"`
	Semester    string `json:"semester"`
	DayofWeek   string `json:"dayofweek"`
	Time        string `json:"time"`
	Teacher     string `json:"teacher"`
	LectureName string `json:"lectureName"`
}

//CallRectures は講義の内容を呼び出す
func CallRectures(grade string, department string, semester string, dayofWeek string, time string, teacher string, lectureName string) ([]Recture, error) {

	//空の構造体のスライスを作成
	Rectures := make([]Recture, 0, 0)

	return Rectures, nil
}
