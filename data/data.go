package data

// UserAuth ユーザー情報
type UserAuth struct {
	ID        int
	Name      string
	Email     string
	Introduce string
	Password  string
}

// UserResponse レスポンスデータ
type UserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Introduce    string `json:"introduce"`
	ErrorMessage string `json:"errorMessage"`
}
