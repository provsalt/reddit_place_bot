package config

type Config struct {
	XAxis int `json:"xAxis"`
	YAxis int `json:"yAxis"`
	Delay int `json:"delay"`

	Accounts []Account `json:"accounts"`
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`

	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
