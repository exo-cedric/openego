package cmd

import (
	"fmt"
	"os"

	exoV2 "github.com/exoscale/egoscale/v2"
	exoApi "github.com/exoscale/egoscale/v2/api"
)

const (
	CLIENT_DEFAULT_ENVIRONMENT = "api"
)

var CLIENT_ENVIRONMENT string

func init() {
	CLIENT_ENVIRONMENT = CLIENT_DEFAULT_ENVIRONMENT
	if v := os.Getenv("EXOSCALE_API_ENVIRONMENT"); v != "" {
		CLIENT_ENVIRONMENT = v
	}
}

// Return a new API client
func NewClient(opts ...exoV2.ClientOpt) (client *exoV2.Client, err error) {
	// Endpoint
	endpoint := exoApi.EndpointURL
	if v := os.Getenv("EXOSCALE_API_ENDPOINT"); v != "" {
		endpoint = v
	}

	// Client
	client, err = exoV2.NewClient(
		os.Getenv("EXOSCALE_API_KEY"),
		os.Getenv("EXOSCALE_API_SECRET"),
		exoV2.ClientOptWithAPIEndpoint(endpoint),
	)

	// Additionial options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("Failed to configrue client: %s", err)
		}
	}

	// Done
	return
}
