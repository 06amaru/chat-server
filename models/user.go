package models

type User struct {
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	PrivateKey string `json:"private,omitempty"`
	PublicKey  []byte `json:"public,omitempty"`
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type HealthCheck struct {
	Status string `json:"healthCheck,omitempty"`
}
