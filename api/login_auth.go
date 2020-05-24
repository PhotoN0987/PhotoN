package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/y-moriwake/PhotoN/api/apiresponse"
	"github.com/y-moriwake/PhotoN/config"
	"github.com/y-moriwake/PhotoN/model/repository"
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

	// 取得
	var err error
	user, err := repository.GetUser(paramUserID, paramPassword)

	// 取得内容精査
	switch err {
	case nil:
		config.Logger.Println("ログイン認証API処理結果:認証成功")
		inResponseStatus(w, http.StatusOK)
	case sql.ErrNoRows:
		config.Logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusUnauthorized)
	default:
		config.Logger.Println("ログイン認証API処理結果:認証失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
	}

	// レスポンスデータ生成
	userInfo := apiresponse.UserResponse{user.ID, user.Name, user.Email, user.Introduce}
	js, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

	defer config.Logger.Println("/api/login  ----------end")
}
