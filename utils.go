package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

// 初期化処理(main()より先に実行)
func init() {

	logfile, _ := os.OpenFile("photoN.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiLogfile)
	logger = log.New(logfile, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// ログレベル設定
func logInfo(args ...interface{}) {
	logger.SetPrefix("INFO:")
	logger.Println(args...)
}

func logError(args ...interface{}) {
	logger.SetPrefix("ERROR:")
	logger.Println(args...)
}

func logWarning(args ...interface{}) {
	logger.SetPrefix("WARNING:")
	logger.Println(args...)
}

// DB接続サンプル
func connDbSample(w http.ResponseWriter, r *http.Request) {

	// データベース接続
	db, err := sql.Open("mysql", "admin:Moriwake1@tcp(moriwake-rds.cfi6vbfhf6p0.ap-northeast-1.rds.amazonaws.com:3306)/photoN?charset=utf8")
	if err != nil {
		logError(err.Error())
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		fmt.Fprintf(w, "データベース接続失敗")
	} else {
		fmt.Fprintf(w, "データベース接続成功")
	}

	// クエリ発行
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	// カラム名を取得
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Fprintf(w, columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

}

// ステータス設定(TODO エラーメッセージ追加)
func inResponseStatus(w http.ResponseWriter, status int) {
	// レスポンスヘッダ追加(WriteHeaderの後に記述すると無効)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
}
