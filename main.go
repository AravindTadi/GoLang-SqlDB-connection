package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result, err := db.Exec("insert into data values(4 ,'xyz')")
	checkError(err)

	lastInsertedId, err := result.LastInsertId()
	fmt.Println("Id :", lastInsertedId)
	checkError(err)

	rowsAffected, err := result.RowsAffected()
	fmt.Println("Rows Affected : ", rowsAffected)

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var d Data
		err := rows.Scan(&d.id, &d.name)

		checkError(err)
		fmt.Println(d)

	}
}
