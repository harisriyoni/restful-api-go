package app

import (
	"database/sql"

	"github.com/harisriyoni/restful-api-go/helper"
)

func NewDB() *sql.DB {

	dsn := "root@tcp(localhost:3306)/belajargo"

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)
	return db
}
