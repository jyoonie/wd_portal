package models

type Login struct {
	EmailAddress string `json:"email_address,omitempty"`
	Password     string `json:"password,omitempty"`
}

type Token struct {
	Token string `json:"token,omitempty"`
}
