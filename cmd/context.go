package cmd

import (
	"context"
	"time"

	exoApi "github.com/exoscale/egoscale/v2/api"
)

const (
	CONTEXT_DEFAULT_TIMEOUT = 60 * time.Second
)

var ROOT_CONTEXT context.Context
var ROOT_CONTEXT_CANCEL context.CancelFunc

func init() {
	ROOT_CONTEXT, ROOT_CONTEXT_CANCEL = context.WithCancel(context.Background())
}

// Use the root context (singleton)
func UseRootContext() (context.Context, context.CancelFunc) {
	return ROOT_CONTEXT, ROOT_CONTEXT_CANCEL
}

// Return a new API context (as per given zone/environment)
func NewApiContext(zone string) (context.Context) {
	timeoutContext, _ := context.WithTimeout(ROOT_CONTEXT, CONTEXT_DEFAULT_TIMEOUT)
	return exoApi.WithEndpoint(timeoutContext, exoApi.NewReqEndpoint(CLIENT_ENVIRONMENT, zone))
}
func NewApiContextWithTimeout(zone string, timeout time.Duration) (context.Context) {
	timeoutContext, _ := context.WithTimeout(ROOT_CONTEXT, timeout)
	return exoApi.WithEndpoint(timeoutContext, exoApi.NewReqEndpoint(CLIENT_ENVIRONMENT, zone))
}
func NewApiContextWithParent(zone string, parent context.Context) (context.Context) {
	return exoApi.WithEndpoint(parent, exoApi.NewReqEndpoint(CLIENT_ENVIRONMENT, zone))
}
