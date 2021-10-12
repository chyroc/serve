package internal

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateEngine(dir string) *gin.Engine {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.NoRoute(func(c *gin.Context) {
		contentType, res, err := ReadPath(dir, c.Request.URL.Path)
		if err != nil {
			c.Data(200, "text/html; charset=utf-8", []byte(Error(err)))
			return
		}
		c.Data(200, contentType+"; charset=utf-8", []byte(res))
	})
	return r
}
