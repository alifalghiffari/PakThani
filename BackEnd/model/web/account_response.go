package web

type AccountResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	NoTelepon string `json:"no_telepon"`
	Role      string `json:"role"`
}
