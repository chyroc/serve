package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "serve",
		Usage:  "Static file hosting",
		Action: runApp,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runApp(c *cli.Context) error {
	return nil
}
