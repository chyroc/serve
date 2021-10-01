package main

import (
	"log"
	"os"

	"github.com/chyroc/serve/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "serve",
		Usage:  "Static file hosting",
		Action: runApp,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "https",
				Aliases: []string{"S"},
				Usage:   "host with https",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runApp(c *cli.Context) error {
	isHttps := c.Bool("https")

	r := internal.CreateEngine()

	return internal.Listen(r, isHttps)
}
