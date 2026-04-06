package dockerhub

type DockerHub struct {
	RegistryProvider string
	RegistryDomain   string
	RegistryUsername string
}

func NewDockerHubProvider(domain, username string) *DockerHub {
	return &DockerHub{
		RegistryProvider: "docker",
		RegistryDomain:   domain,
		RegistryUsername: username,
	}
}

func (s *DockerHub) Username() string { return s.RegistryUsername }
func (s *DockerHub) Provider() string { return s.RegistryProvider }
func (s *DockerHub) Domain() string   { return s.RegistryDomain }
