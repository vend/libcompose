package app

import (
	"github.com/vend/libcompose/cli/logger"
	"github.com/vend/libcompose/docker"
	"github.com/vend/libcompose/docker/ctx"
	"github.com/vend/libcompose/project"
	"github.com/urfave/cli"
)

// ProjectFactory is a struct that holds the app.ProjectFactory implementation.
type ProjectFactory struct {
}

// Create implements ProjectFactory.Create using docker client.
func (p *ProjectFactory) Create(c *cli.Context) (project.APIProject, error) {
	context := &ctx.Context{}
	context.LoggerFactory = logger.NewColorLoggerFactory()
	Populate(context, c)
	return docker.NewProject(context, nil)
}
