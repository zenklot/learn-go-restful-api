package app

import (
	"database/sql"
	"time"

	"github.com/zenklot/learn-go-restful-api/helper"
)

func NewDB() *sql.DB {
	db, er := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError(er)

	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 10)
	return db
}
