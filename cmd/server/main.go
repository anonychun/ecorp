package main

import (
	"context"
	"log"
	"os"

	"github.com/anonychun/ecorp/internal/server"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{}

	cmd.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "Start the server",
			Action: func(ctx context.Context, c *cli.Command) error {
				return server.Start(ctx)
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to run command:", err)
	}
}
