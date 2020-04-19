package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// メイン処理
func main() {

	logger.Println("----------サーバー起動開始----------")

	mux := http.NewServeMux()

	// ハンドラ
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api/login", loginHandlar)
	mux.HandleFunc("/api/signup", signUpHandlar)

	// サーバー設定
	server := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラー
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
