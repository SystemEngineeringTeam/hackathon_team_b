package db

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type ReviewDetail struct{
// 	ReviewStar int `json:"department"`
// 	contents string `json:"contents"`
// }



// //CallReview はレビューの構造体のスライスを返す
// func CallReview(indexNumber int)  {


// 	db, err := sql.Open("mysql", "root@/with_b")
//     if err != nil {
//         log.Println(err)
//     }
// 	defer db.Close()


// 	rows,err:=db.Query("SELECT * FROM reviews WHERE indexNumber=?",indexNumber)
// 	defer rows.Close()
//     if err != nil {
//         log.Println(err)
//     }

// 	 ReviewDetails:=make([]ReviewDetail,0,0)


// 	for rows.Next(){

// 		var star int
// 		var contents string


// 		if err:=rows.Scan(&star,&contents);err!=nil{
// 			log.Print(err)
// 		}
// 		tmpReview:=ReviewDetails{}

// 	}



// }







