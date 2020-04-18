package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/ini.v1"
)

type DbConfig struct {
	User string
	Pass string
	Host string
	Name string
}

var logger *log.Logger
var Cnf DbConfig

// 初期化処理(main()より先に実行)
func init() {

	logfile, _ := os.OpenFile("photoN.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiLogfile)
	logger = log.New(logfile, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	c, _ := ini.Load("./config.ini")
	Cnf = DbConfig{
		User: c.Section("db").Key("user").String(),
		Pass: c.Section("db").Key("pass").String(),
		Host: c.Section("db").Key("host").String(),
		Name: c.Section("db").Key("name").String(),
	}
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
