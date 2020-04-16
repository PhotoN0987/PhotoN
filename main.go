package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// メイン処理
func main() {

	logInfo("----------サーバー起動開始----------")

	mux := http.NewServeMux()

	// ハンドラ
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api/login", loginHandlar)

	// サーバー設定
	server := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラー
func index(w http.ResponseWriter, r *http.Request) {

}
