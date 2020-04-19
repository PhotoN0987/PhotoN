package main

import (
	"database/sql"
	"log"
	"net/http"
)

//User is struct
type User struct {
	email    string
	password string
}

// ログイン機能
func loginHandlar(w http.ResponseWriter, r *http.Request) {

	logInfo("/api/login  ----------start")

	// リクエストパラメータ
	paramUserID := r.FormValue("userId")
	paramPassword := r.FormValue("password")
	logInfo("userId：    ", paramUserID)
	logInfo("password：  ", paramPassword)

	// パラメータチェック
	if paramUserID == "" || paramPassword == "" {
		inResponseStatus(w, http.StatusBadRequest)
		return
	}

	// DB接続
	db, err := sql.Open("mysql", Cnf.User+":"+Cnf.Pass+"@"+Cnf.Host+"/"+Cnf.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// クエリ定義
	var user User
	query := `
		SELECT user_id,user_password 
		FROM users
		WHERE users.user_email = ? and users.user_password = ?`

	// クエリ実行&マッピング
	err = db.QueryRow(query, paramUserID, paramPassword).Scan(&user.email, &user.password)
	logInfo("取得結果:user_id=", user.email, " password=", user.password)

	// 取得内容精査
	if err == sql.ErrNoRows {
		logError("ユーザ認証失敗")
		inResponseStatus(w, http.StatusUnauthorized)
	} else {
		logInfo("ユーザー認証成功")
		inResponseStatus(w, http.StatusOK)
	}

	defer logInfo("/api/login  ----------end")
}
