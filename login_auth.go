package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/y-moriwake/PhotoN/data"
)

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
	var user data.UserAuth
	query := `
		SELECT user_id,user_name,user_email,user_password,user_introduce
		FROM users
		WHERE users.user_email = ? and users.user_password = ?`

	// クエリ実行&マッピング
	err = db.QueryRow(query, paramUserID, paramPassword).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Introduce)
	logger.Println("取得結果:user_id =", user.ID, ",user_name =", user.Name, ",user_email =", user.Email, ",user_password =", user.Password, ",user_introduce =", user.Introduce)

	// 取得内容精査
	var errorMessage string
	switch err {
	case nil:
		logger.Println("ログイン認証API処理結果:認証成功")
		inResponseStatus(w, http.StatusOK)
	case sql.ErrNoRows:
		logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusUnauthorized)
		errorMessage = err.Error()
	default:
		logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		errorMessage = err.Error()
	}

	// レスポンスデータ生成
	userInfo := data.UserResponse{user.ID, user.Name, user.Email, user.Introduce, errorMessage}
	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

	defer logger.Println("/api/login  ----------end")
}
