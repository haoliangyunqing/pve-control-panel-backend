package models

type PVELoginRequest struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Realm    string `json:"realm"`
}
type PVEAuthResponse struct {
	Data struct {
		Ticket string `json:"ticket"`              // 认证票据
		CSRF   string `json:"CSRFPreventionToken"` // CSRF令牌
	} `json:"data"`
}
