package auth

type SignInRequest struct {
	IpAddress    string `json:"-"`
	UserAgent    string `json:"-"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
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
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
	} `json:"admin"`
}
