package model

import (
	"database/sql"
	"log"

	"github.com/y-moriwake/PhotoN/config"
)

// DbConfig DB設定の構造体
type DbConfig struct {
	User string
	Pass string
	Host string
	Name string
}

// DB接続
func dbConnect() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.Cnf.User+":"+config.Cnf.Pass+"@"+config.Cnf.Host+"/"+config.Cnf.Name)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
