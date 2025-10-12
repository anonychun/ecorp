package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/anonychun/benih/cmd/generate/internal"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{}

	cmd.Commands = []*cli.Command{
		{
			Name:  "migration",
			Usage: "Generate a new database migration",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name: "name",
				},
				&cli.StringArg{
					Name:  "type",
					Value: "sql",
				},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				name := c.StringArg("name")
				if name == "" {
					return errors.New("missing migration name")
				}

				return internal.GenerateMigration(name, c.StringArg("type"))
			},
		},
		{
			Name:  "app",
			Usage: "Generate a new app",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name: "name",
				},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				name := c.StringArg("name")
				if name == "" {
					return errors.New("missing app name")
				}

				return internal.GenerateApp(name)
			},
		},
		{
			Name:  "repository",
			Usage: "Generate a new repository",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name: "name",
				},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				name := c.StringArg("name")
				if name == "" {
					return errors.New("missing repository name")
				}

				return internal.GenerateRepository(name)
			},
		},
		{
			Name:  "entity",
			Usage: "Generate a new entity",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name: "name",
				},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				name := c.StringArg("name")
				if name == "" {
					return errors.New("missing entity name")
				}

				return internal.GenerateEntity(name)
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to run command:", err)
	}
}
