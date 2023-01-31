package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "bad"
	dbPass := "123qwe"
	dbName := "crm_vcgasia"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
