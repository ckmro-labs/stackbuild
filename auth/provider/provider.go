package provider

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

// ProviderNameInitializationMap is a map of initialization functions for supported OAUTH providers
var ProviderNameInitializationMap = make(map[string]func(key, secret, cbUrl string, scopes ...string) goth.Provider)

func init() {
	ProviderNameInitializationMap["github"] = func(key, secret, cbUrl string, scopes ...string) goth.Provider {
		provider := github.New(key, secret, cbUrl, scopes...)
		return provider
	}
}

// Provider represents an OAUTH provider for mouthful
type Provider struct {
	Name           string
	secret         string
	key            string
	Implementation *goth.Provider
}

// New returns a new provider with the given parameters. It checks if the provider is supported or not and if all the requierements are met.
func New(name string, secret, key string, uri string) (*Provider, error) {
	var initfunction func(key, secret, cbUrl string, scopes ...string) goth.Provider
	if val, ok := ProviderNameInitializationMap[name]; !ok {
		return nil, fmt.Errorf("No such OAUTH provider %v", name)
	} else {
		initfunction = val
	}

	if uri == "" {
		return nil, fmt.Errorf("Invalid callback uri provided for OAUTH provider %v", name)
	}
	gothProvider := initfunction(key, secret, uri)

	return &Provider{
		Implementation: &gothProvider,
		Name:           name,
		secret:         secret,
		key:            key,
	}, nil
}
