package web

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	NoTelepon string `json:"no_telepon"`
	Role      string `json:"role"`
	Token     string `json:"token"`
	Error     string `json:"error"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
