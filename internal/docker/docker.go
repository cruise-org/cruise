package docker

import (
	"log"

	"github.com/docker/docker/client"
)

var cli *client.Client

func init() {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("Docker Error: " + err.Error())
	}

	cli = c
}

func GetClient() *client.Client {
	return cli
}
