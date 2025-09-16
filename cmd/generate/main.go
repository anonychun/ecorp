package main

import (
	"context"
	"log"
	"os"

	"github.com/anonychun/ecorp/cmd/generate/internal"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{}

	cmd.Commands = []*cli.Command{
		{
			Name:  "migration",
			Usage: "Generate a new database migration",
			Action: func(ctx context.Context, c *cli.Command) error {
				return internal.GenerateMigration(c.Args().Get(0), c.Args().Get(1))
			},
		},
		{
			Name:  "app",
			Usage: "Generate a new app",
			Action: func(ctx context.Context, c *cli.Command) error {
				return nil
			},
		},
		{
			Name:  "repository",
			Usage: "Generate a new repository",
			Action: func(ctx context.Context, c *cli.Command) error {
				return nil
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to run command:", err)
	}
}
