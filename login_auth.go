package main

import (
	"database/sql"
	"net/http"
)

//User is struct
type User struct {
	id       int
	email    string
	password string
}

// ログイン機能
func loginHandlar(w http.ResponseWriter, r *http.Request) {

	logger.Println("/api/login  ----------start")

	// リクエストパラメータ
	paramUserID := r.FormValue("userId")
	paramPassword := r.FormValue("password")
	logger.Println("パラメータ:userId =", paramUserID)
	logger.Println("パラメータ:password =", paramPassword)

	// パラメータチェック
	if paramUserID == "" || paramPassword == "" {
		inResponseStatus(w, http.StatusBadRequest)
		return
	}

	// データベース接続
	db, err := dbConnect()
	defer db.Close()

	// クエリ定義
	var user User
	query := `
		SELECT user_id,user_email,user_password 
		FROM users
		WHERE users.user_email = ? and users.user_password = ?`

	// クエリ実行&マッピング
	err = db.QueryRow(query, paramUserID, paramPassword).Scan(&user.id, &user.email, &user.password)
	logger.Println("取得結果:user_id =", user.id, ",user_email =", user.email, ",user_password =", user.password)

	// 取得内容精査
	if err == nil {
		logger.Println("ログイン認証API処理結果:認証成功")
		inResponseStatus(w, http.StatusOK)
	} else if err == sql.ErrNoRows {
		logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusUnauthorized)
	} else {
		logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
	}

	defer logger.Println("/api/login  ----------end")
}
