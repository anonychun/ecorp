package auth

type LoginRequest struct {
	IpAddress    string `json:"-"`
	UserAgent    string `json:"-"`
	EmailAddress string `json:"emailAddress" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"-"`
	Admin struct {
		Id           string `json:"id"`
		EmailAddress string `json:"emailAddress"`
	} `json:"admin"`
}

type LogoutRequest struct {
	Token string
}

type MeResponse struct {
	Admin struct {
		Id           string `json:"id"`
		EmailAddress string `json:"emailAddress"`
	} `json:"admin"`
}
