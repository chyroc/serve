package main

import (
	"fmt"
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
	dir := c.Args().First()
	isHttps := c.Bool("https")

	fmt.Printf("dir=%q, https=%v\n", dir, isHttps)

	r := internal.CreateEngine(dir)

	return internal.Listen(r, isHttps)
}
