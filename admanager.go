//go:generate ./scripts/generate.sh

package admanager

import (
	"context"
	"net/http"

	"github.com/hooklift/gowsdl/soap"
	"golang.org/x/oauth2"
)

// Config holds the network-level settings for the Ad Manager API.
type Config struct {
	NetworkCode     string
	ApplicationName string
}

// Client is a version-agnostic Ad Manager API client. It holds the
// authenticated HTTP client and configuration. Use a version-specific
// NewService() helper (e.g. v202505.NewService) to obtain a *soap.Client
// bound to a particular API version.
type Client struct {
	Config     Config
	HTTPClient *http.Client
}

// NewClient creates a Client with an HTTP client that automatically attaches
// and refreshes OAuth2 Bearer tokens via the provided TokenSource.
func NewClient(ctx context.Context, cfg Config, ts oauth2.TokenSource) *Client {
	return &Client{
		Config:     cfg,
		HTTPClient: oauth2.NewClient(ctx, ts),
	}
}

// NewServiceClient creates a *soap.Client for the given API version and service
// name, pre-configured with OAuth2 authentication and the SOAP RequestHeader.
// This is intended to be called by per-version helper packages (e.g.
// services/v202505) rather than directly by end users.
func NewServiceClient(c *Client, version, serviceName string) *soap.Client {
	url := "https://ads.google.com/apis/ads/publisher/" + version + "/" + serviceName
	soapClient := soap.NewClient(url, soap.WithHTTPClient(c.HTTPClient))
	soapClient.AddHeader(NewRequestHeader(version, c.Config.NetworkCode, c.Config.ApplicationName))
	return soapClient
}
