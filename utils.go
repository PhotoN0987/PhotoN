package main

import (
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

// ステータス設定(TODO エラーメッセージ追加)
func inResponseStatus(w http.ResponseWriter, status int) {
	// レスポンスヘッダ追加(WriteHeaderの後に記述すると無効)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
}
