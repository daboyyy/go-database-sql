package main

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Cover struct {
	Id int
	Name string
}

var db *sqlx.DB

func main() {
	var err error

	// MSSQL
	// db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?database=techcoach")

	// MySQL
	db, err = sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73)/techcoach")

	if err != nil {
		panic(err)
	}

	// GetCovers
	covers, err := GetCovers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, cover := range covers {
		fmt.Println(cover)
	}
}

func GetCovers() ([]Cover, error) {
	query := "select id, name from cover"
	covers := []Cover{}
	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}

	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	query := "select id, name from cover where id=?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}

	return &cover, nil
}
