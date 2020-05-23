package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/y-moriwake/PhotoN/config"
)

// ステータス設定
func inResponseStatus(w http.ResponseWriter, status int) {
	// レスポンスヘッダ追加(WriteHeaderの後に記述すると無効)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// DB接続
func dbConnect() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.Cnf.User+":"+config.Cnf.Pass+"@"+config.Cnf.Host+"/"+config.Cnf.Name)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
