package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "quynhnhu2010"
	dbName   = "ieltscenter"
	db       *sql.DB
	tx       *sql.Tx
	err error
)

func ConnectDB() {

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}

func GetTx() *sql.Tx {

	tx, err = db.Begin()
	if err != nil {
		panic(err)
	}
	return tx
}