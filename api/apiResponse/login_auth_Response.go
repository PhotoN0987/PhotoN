package apiresponse

// UserResponse レスポンスデータ
type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Introduce string `json:"introduce"`
}
