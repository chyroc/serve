package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chyroc/serve/internal"
	"github.com/gin-gonic/gin"
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
	dir := c.Args().First()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.NoRoute(func(c *gin.Context) {
		isHtml, res, err := internal.ReadPath(dir, c.Request.URL.Path)
		if err != nil {
			c.Data(200, "text/html; charset=utf-8", []byte(internal.Error(err)))
			return
		}
		if isHtml {
			c.Data(200, "text/html; charset=utf-8", []byte(res))
			return
		}
		c.Data(200, "text/plain; charset=utf-8", []byte(res))
	})

	listener, err := internal.GetAvailableAddress()
	if err != nil {
		return err
	}
	fmt.Println("File Hosting: " + dir)
	fmt.Println("Serve Listening: http://" + listener.Addr().String())
	return r.RunListener(listener)
}
