package models

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private"`
	PublicKey  []byte `json:"public"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
