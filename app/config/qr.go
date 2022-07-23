package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "quynhnhu2010"
	dbName := "ieltscenter"
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}