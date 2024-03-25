package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
	httpClient   *http.Client
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		httpClient:   &http.Client{},
	}
}

func (c *Client) Authenticate(code string, redirectURI string) (string, error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURI},
	}

	resp, err := c.httpClient.PostForm("https://api.mercadolibre.com/oauth/token", data)
	if err != nil {
		return "", fmt.Errorf("error making authentication request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authentication request failed with status code %d", resp.StatusCode)
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("error decoding token response: %w", err)
	}

	c.AccessToken = tokenResponse.AccessToken
	c.RefreshToken = tokenResponse.RefreshToken
	return c.AccessToken, nil
}

// Métodos para as operações HTTP (GET, POST, PUT, DELETE) serão adicionados aqui.
