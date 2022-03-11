package main

import (
	"errors"
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

// begin tran process example
func AddCover(cover Cover) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// MySQL
	query := "insert into cover (id, name) values (?, ?)" // no auto increment have to add id
	result, err := tx.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
