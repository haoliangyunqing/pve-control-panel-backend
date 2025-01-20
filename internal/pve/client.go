package pve

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"pve-control-panel-backend/internal/models"
)

type PVEConfig struct {
	Host     string
	Username string
	Password string
	Realm    string
}

type PVEClient struct {
	client *resty.Client
	config *PVEConfig
	ticket string
	csrf   string
}

func NewPVEClient(cfg *PVEConfig) *PVEClient {
	return &PVEClient{
		client: resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
		config: cfg,
	}
}

func (c *PVEClient) Authenticate() error {
	url := fmt.Sprintf("%s/api2/json/access/ticket", c.config.Host)
	var authResp models.PVEAuthResponse
	resp, err := c.client.R().
		SetFormData(map[string]string{
			"username": c.config.Username,
			"password": c.config.Password,
			"realm":    c.config.Realm,
		}).
		SetResult(&authResp).
		Post(url)
	if err != nil {

		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("auth error: %s", resp.StatusCode())
	}

	c.ticket = authResp.Data.Ticket
	c.csrf = authResp.Data.CSRF
	return nil
}

func (c *PVEClient) GetTicket() string {
	return c.ticket
}

func (c *PVEClient) GetCSRFToken() string {
	return c.csrf
}
