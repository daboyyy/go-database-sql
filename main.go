package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type Cover struct {
	Id int
	Name string
}

var db *sql.DB

func main() {
	var err error

	// MSSQL
	// db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?database=techcoach")

	// MySQL
	db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73)/techcoach")

	if err != nil {
		panic(err)
	}

	// AddCover
	/* cover := Cover{9, "Bond"}
	err = AddCover(cover)
	if err != nil {
		panic(err)
	} */

	// UpdateCover
	/* cover := Cover{9, "Hello"}
	err = UpdateCover(cover)
	if err != nil {
		panic(err)
	} */

	// DeleteCover
	deleteId := 9
	err = DelteCover(deleteId)
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

	// GetCover
	/* cover, err := GetCover(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(cover) */
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
	 	return nil, err
	}

	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // promise if program end this will be close

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}

		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}

		covers = append(covers, cover)
	}

	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// MSSQL Server
	// query := "select id, name from cover where id=@id"

	// MySQL
	query := "select id, name from cover where id=?"

	row := db.QueryRow(query, id)
	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}

	return &cover, nil
}

func AddCover(cover Cover) error {
	// MySQL
	query := "insert into cover (id, name) values (?, ?)" // no auto increment have to add id
	result, err := db.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil
}

func UpdateCover(cover Cover) error {
	// MySQL
	query := "update cover set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil
}

func DelteCover(id int) error {
	// MySQL
	query := "delete from cover where id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}
