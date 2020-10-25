package options

import (
	"io/ioutil"
	"os"

	"github.com/ErickMaria/envcontainer/internal/config"
)

const (
	PATH_DEFAULT   = ".envcontainer/compose/env"
	DOCKERFILE     = ".envcontainer/Dockerfile"
	DOCKER_COMPOSE = ".envcontainer/compose/docker-compose.yaml"
	ENV            = ".envcontainer/compose/env/.env"
)

type Command struct {
}

func (c *Command) Init(flags map[string]string) string {

	if flags["project"] == "" {
		flags["project"] = "envcontainer"
	}

	dc := config.DockerCompose{
		Version: "3.6",
		Services: config.Services{
			Environment: config.Environment{
				Volumes: []config.Volumes{
					config.Volumes{
						Type:   "bind",
						Source: "../../",
						Target: "/home/" + flags["project"],
					},
				},
				Build: config.Build{
					Dockerfile: "Dockerfile",
					Context:    "../",
				},
				Ports: []string{
					flags["listener"],
				},
				WorkingDir: "/home/" + flags["project"],
				EnvFile: []string{
					flags["envfile"],
				},
				StdinOpen: true,
				Tty:       true,
			},
		},
	}

	data := dc.Marshal()

	err := os.MkdirAll(PATH_DEFAULT, 0755)
	check(err)

	createFile := func(name string, data []byte) {
		check(ioutil.WriteFile(name, data, 0644))
	}

	createFile(DOCKERFILE, []byte(""))
	createFile(DOCKER_COMPOSE, data)
	createFile(ENV, []byte(""))

	// archive.FileWrite(DOCKER_COMPOSE, "sfgf")

	return "init"
}

func (c *Command) Up() string {
	return "up"
}

func (c *Command) Build() string {
	return "build"
}

func (c *Command) Down() string {
	return "down"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
