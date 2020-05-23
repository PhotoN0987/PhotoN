package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/y-moriwake/PhotoN/config"
	"github.com/y-moriwake/PhotoN/model"
)

// LoginHandlar ログイン機能
func LoginHandlar(w http.ResponseWriter, r *http.Request) {

	config.Logger.Println("/api/login  ----------start")

	// リクエストパラメータ
	paramUserID := r.FormValue("userId")
	paramPassword := r.FormValue("password")
	config.Logger.Println("パラメータ:userId =", paramUserID)
	config.Logger.Println("パラメータ:password =", paramPassword)

	// パラメータチェック
	if paramUserID == "" || paramPassword == "" {
		inResponseStatus(w, http.StatusBadRequest)
		return
	}

	// データベース接続
	db, err := dbConnect()
	defer db.Close()

	// クエリ定義
	var user model.UserAuth
	query := `
		SELECT user_id,user_name,user_email,user_password,user_introduce
		FROM users
		WHERE users.user_email = ? and users.user_password = ?`

	// クエリ実行&マッピング
	err = db.QueryRow(query, paramUserID, paramPassword).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Introduce)
	config.Logger.Println("取得結果:user_id =", user.ID, ",user_name =", user.Name, ",user_email =", user.Email, ",user_password =", user.Password, ",user_introduce =", user.Introduce)

	// 取得内容精査
	var errorMessage string
	switch err {
	case nil:
		config.Logger.Println("ログイン認証API処理結果:認証成功")
		inResponseStatus(w, http.StatusOK)
	case sql.ErrNoRows:
		config.Logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusUnauthorized)
		errorMessage = err.Error()
	default:
		config.Logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		errorMessage = err.Error()
	}

	// レスポンスデータ生成
	userInfo := model.UserResponse{user.ID, user.Name, user.Email, user.Introduce, errorMessage}
	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

	defer config.Logger.Println("/api/login  ----------end")
}
