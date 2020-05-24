package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/y-moriwake/PhotoN/api/apirequest"
	"github.com/y-moriwake/PhotoN/config"
)

// SignUpHandlar ユーザー新規登録
func SignUpHandlar(w http.ResponseWriter, r *http.Request) {

	config.Logger.Println("/api/signup  ----------start")

	// パラメータ解析
	body, _ := ioutil.ReadAll(r.Body)
	var postedBody apirequest.SignUpRequestBody
	json.Unmarshal(body, &postedBody)

	userName := postedBody.Name
	userEmail := postedBody.Email
	userPasseord := postedBody.Password

	//　パラメータチェック
	config.Logger.Println("パラメータ:name =", userName)
	config.Logger.Println("パラメータ:email =", userEmail)
	config.Logger.Println("パラメータ:password =", userPasseord)
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
		config.Logger.Println("ユーザー新規登録API:登録失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		return
	}
	defer ins.Close()

	// 登録実行
	_, err = ins.Exec(userPasseord, userName, userEmail)
	if err != nil {
		config.Logger.Println("ユーザー新規登録API:登録失敗", err)
		inResponseStatus(w, http.StatusInternalServerError)
		return
	}

	// 成功した場合
	config.Logger.Println("ユーザー新規登録API:登録成功")
	inResponseStatus(w, http.StatusOK)

	defer config.Logger.Println("/api/signup  ----------end")
}
