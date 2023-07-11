package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hongminhcbg/base-image/src/liveness"
	"github.com/urfave/cli/v3"
)

var app = new(cli.Command)

func hello(ctx *cli.Context) error {
	fmt.Println("hello world")
	time.Sleep(24 * time.Hour)
	return nil
}

func main() {
	app.Commands = []*cli.Command{
		{
			Name:        "hello",
			Usage:       "hello world",
			Description: "hello world",
			Action:      hello,
		},
		{
			Name:        "liveness",
			Usage:       "run liveness server",
			Description: "run liveness server",
			Action:      liveness.Main,
		},
		{
			Name:        "liveness_dep",
			Usage:       "run liveness server with dependencies health",
			Description: "run liveness server with dependencies health",
			Action:      liveness.LivenessDependencies,
		},
	}

	app.Run(context.Background(), os.Args)
}
