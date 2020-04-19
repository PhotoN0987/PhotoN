package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SignUpRequestBody struct {
	Name     string
	Email    string
	Password string
}

// ユーザー新規登録
func signUpHandlar(w http.ResponseWriter, r *http.Request) {

	logInfo("/api/signup  ----------start")

	// パラメータ解析
	body, _ := ioutil.ReadAll(r.Body)
	var postedBody SignUpRequestBody
	json.Unmarshal(body, &postedBody)

	userName := postedBody.Name
	userEmail := postedBody.Email
	userPasseord := postedBody.Password

	//　パラメータチェック
	logInfo("name:", userName)
	logInfo("email:", userEmail)
	logInfo("password:", userPasseord)
	if userName == "" || userEmail == "" || userPasseord == "" {
		inResponseStatus(w, http.StatusBadRequest)
		return
	}

	// DB接続
	db, err := sql.Open("mysql", Cnf.User+":"+Cnf.Pass+"@"+Cnf.Host+"/"+Cnf.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 登録
	query := `
		INSERT INTO users
			(user_id, user_password, user_name, user_email, user_introduce, delete_flag) 
		VALUES ("",?,?,?,"",0)`

	ins, err := db.Prepare(query)
	if err != nil {
		log.Fatal("ユーザー新規登録失敗", err)
	}
	ins.Exec(userPasseord, userName, userEmail)

	inResponseStatus(w, http.StatusOK)

	defer logInfo("/api/signup  ----------end")
}
