package api

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func NewConfig(clientID, clientSecret, redirectURI string) *Config {
	return &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}
}
