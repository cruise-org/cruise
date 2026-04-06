package registry

import (
	"fmt"

	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/registry/dockerhub"
	"github.com/zalando/go-keyring"
)

type Registry interface {
	Username() string
	Provider() string // "Harbor", "DockerHub"
	Domain() string   // "harbor.local", "docker.io"

	Login(password string) error
	Logout() error

	// Listing
	// ListImages() ([]types.RegistryImage, error)
	// TODO: Add these
	// ImageDetails() ([]types.RegistryImage, error)

	// Operations
	// Pull(image types.RegistryImage) error
	// Push(image types.RegistryImage) error
	// Retag(image types.RegistryImage, newTag string) error
}

func GetRegistry(r *config.RegistryConfig) (Registry, error) {
	switch r.Provider {
	case "dockerhub":
		return dockerhub.NewDockerHubProvider(r.Domain, r.Username), nil
	default:
		return nil, fmt.Errorf("provider not supported, received: %s", r.Provider)
	}
}

func IsLoggedIn(user string) bool {
	_, err := keyring.Get("cruise", user)
	return err == nil
}
