package lyft

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// OAuth 2.0 Endpoint URIs
	AuthURL  = "https://api.lyft.com/oauth/authorize"
	TokenURL = "https://api.lyft.com/oauth/token"

	// OAuth 2.0 Scopes
	Offline      = "offline"
	Profile      = "profile"
	Public       = "public"
	RidesRead    = "rides.read"
	RidesRequest = "rides.request"
)

func Endpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  AuthURL,
		TokenURL: TokenURL}
}

func NewClientCredentials(ctx context.Context, clientId, clientSecret string, scopes []string) *http.Client {
	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     TokenURL,
		Scopes:       scopes}

	return config.Client(ctx)
}
