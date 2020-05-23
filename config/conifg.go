package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// DbConfig DB設定の構造体
type DbConfig struct {
	User string
	Pass string
	Host string
	Name string
}

// Logger ロガークラス
var Logger *log.Logger

// Cnf データベースコンフィグ
var Cnf DbConfig

// 初期化処理(main()より先に実行)
func init() {

	// ログ設定
	logfile, _ := os.OpenFile("photoN.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiLogfile)
	Logger = log.New(logfile, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	// データベース設定読み込み
	c, _ := ini.Load("./config.ini")
	Cnf = DbConfig{
		User: c.Section("db").Key("user").String(),
		Pass: c.Section("db").Key("pass").String(),
		Host: c.Section("db").Key("host").String(),
		Name: c.Section("db").Key("name").String(),
	}
}
