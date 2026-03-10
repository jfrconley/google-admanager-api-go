package admanager

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scope is the OAuth2 scope required for the Google Ad Manager SOAP API.
const Scope = "https://www.googleapis.com/auth/dfp"

// ServiceAccountTokenSourceFromJSON returns an oauth2.TokenSource for a Google
// service account using the provided JSON key bytes.
func ServiceAccountTokenSourceFromJSON(ctx context.Context, jsonKey []byte) (oauth2.TokenSource, error) {
	creds, err := google.CredentialsFromJSON(ctx, jsonKey, Scope)
	if err != nil {
		return nil, err
	}
	return creds.TokenSource, nil
}

// ServiceAccountTokenSourceFromFile returns an oauth2.TokenSource for a Google
// service account using the JSON key file at the given path.
func ServiceAccountTokenSourceFromFile(ctx context.Context, keyFile string) (oauth2.TokenSource, error) {
	data, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return ServiceAccountTokenSourceFromJSON(ctx, data)
}

// TokenSourceFromRefreshToken returns an oauth2.TokenSource that uses the
// provided OAuth2 client credentials and refresh token to obtain access tokens.
func TokenSourceFromRefreshToken(ctx context.Context, clientID, clientSecret, refreshToken string) oauth2.TokenSource {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{Scope},
	}
	return cfg.TokenSource(ctx, &oauth2.Token{RefreshToken: refreshToken})
}

// StaticTokenSource returns an oauth2.TokenSource that always returns the
// provided access token. Useful for short-lived scripts where the caller
// manages token refresh externally.
func StaticTokenSource(accessToken string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
}
