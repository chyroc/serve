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
			&cli.BoolFlag{
				Name:  "disable-cors",
				Usage: "disable cors",
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
	isDisableCors := c.Bool("disable-cors")

	fmt.Printf("dir=%q, https=%v, isDisableCors=%v\n", dir, isHttps, isDisableCors)

	r := internal.CreateEngine(dir, isDisableCors)

	return internal.Listen(r, isHttps)
}
