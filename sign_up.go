package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SignUpRequestBody リクエストのJSON
type SignUpRequestBody struct {
	Name     string
	Email    string
	Password string
}

// ユーザー新規登録
func signUpHandlar(w http.ResponseWriter, r *http.Request) {

	logger.Println("/api/signup  ----------start")

	// パラメータ解析
	body, _ := ioutil.ReadAll(r.Body)
	var postedBody SignUpRequestBody
	json.Unmarshal(body, &postedBody)

	userName := postedBody.Name
	userEmail := postedBody.Email
	userPasseord := postedBody.Password

	//　パラメータチェック
	logger.Println("パラメータ:name =", userName)
	logger.Println("パラメータ:email =", userEmail)
	logger.Println("パラメータ:password =", userPasseord)
	if userName == "" || userEmail == "" || userPasseord == "" {
		inResponseStatus(w, http.StatusBadRequest)
		return
	}

	// DB接続
	db, err := dbConnect()
	defer db.Close()

	// 登録
	query := `
		INSERT INTO users
			(user_id, user_password, user_name, user_email, user_introduce, delete_flag) 
		VALUES ("",?,?,?,"",0)`

	ins, err := db.Prepare(query)
	if err != nil {
		logger.Println("ーザー新規登録API:登録失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		return
	}
	defer ins.Close()

	_, err = ins.Exec(userPasseord, userName, userEmail)
	if err != nil {
		logger.Println("ーザー新規登録API:登録失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		return
	}

	logger.Println("ユーザー新規登録API:登録成功")
	inResponseStatus(w, http.StatusOK)

	defer logger.Println("/api/signup  ----------end")
}
