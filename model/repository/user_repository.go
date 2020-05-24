package repository

import (
	"github.com/y-moriwake/PhotoN/config"
	"github.com/y-moriwake/PhotoN/model"
	"github.com/y-moriwake/PhotoN/model/data"
)

// GetUser 一件取得する
func GetUser(email string, password string) (data.User, error) {

	// クエリ定義
	var user data.User
	query := `
			SELECT user_id,user_name,user_email,user_password,user_introduce
			FROM users
			WHERE users.user_email = ? and users.user_password = ?`

	// データベース接続
	db, err := model.DbConnect()
	defer db.Close()

	// クエリ実行&マッピング
	err = db.QueryRow(query, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Introduce)
	config.Logger.Println("取得結果:user_id =", user.ID, ",user_name =", user.Name, ",user_email =", user.Email, ",user_password =", user.Password, ",user_introduce =", user.Introduce)

	return user, err
}
