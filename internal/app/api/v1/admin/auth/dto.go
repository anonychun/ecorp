package auth

type SignInRequest struct {
	IpAddress    string `json:"-"`
	UserAgent    string `json:"-"`
	EmailAddress string `json:"emailAddress" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string
}

type SignOutRequest struct {
	Token string
}

type MeResponse struct {
	Admin struct {
		Id           string `json:"id"`
		EmailAddress string `json:"emailAddress"`
	} `json:"admin"`
}
