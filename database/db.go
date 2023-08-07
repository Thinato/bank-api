package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
    db *sql.DB
}

func NewDatabase() (*Database, error) {
    data_source_name := "root:root@tcp(127.0.0.1:3306)/mydb"

    db, err := sql.Open("mysql", data_source_name)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    fmt.Println("Connected to the database!")

    return &Database{
        db: db,
    }, nil
}

func (d *Database ) Exec(query string) (error) {
    _, err := d.db.Exec(query)
    if err != nil{
        return err
    }
    return nil
}

func (d *Database) Close() error {
    err := d.db.Close()
    if err != nil {
        return err
    }
    return nil
}