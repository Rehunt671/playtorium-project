package dtos

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
