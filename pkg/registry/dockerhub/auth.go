package dockerhub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/zalando/go-keyring"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func (s *DockerHub) Login(pass string) error {
	token, err := s.verify(pass)
	if err != nil {
		return err
	}

	return keyring.Set("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()), token)
}

func (s *DockerHub) Logout() error {
	return keyring.Delete("cruise", fmt.Sprintf("%s/%s", s.Provider(), s.Domain()))
}

func (s *DockerHub) verify(secret string) (string, error) {
	client := resty.New()
	var tokenResp tokenResponse

	resp, err := client.R().
		SetBasicAuth(s.Username(), secret).
		SetQueryParams(map[string]string{
			"service": "registry.docker.io",
			"scope":   "repository:library/alpine:pull",
		}).
		SetResult(&tokenResp).
		Get("https://auth.docker.io/token")
	if err != nil {
		log.Printf("[DHUB VERIFY] request failed: %v", err)
		return "", err
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		log.Printf("[DHUB VERIFY] credentials verified for user: %s", s.Username())
		return tokenResp.Token, nil
	case http.StatusUnauthorized:
		log.Printf("[DHUB VERIFY] invalid credentials for user: %s", s.Username())
		return "", fmt.Errorf("invalid credentials")
	default:
		log.Printf("[DHUB VERIFY] unexpected status %d for user: %s", resp.StatusCode(), s.Username())
		return "", fmt.Errorf("unexpected status %d: %s", resp.StatusCode(), resp.String())
	}
}
