package internal

import (
	"github.com/gin-gonic/gin"
)

func CreateEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.NoRoute(func(c *gin.Context) {
		isHtml, res, err := ReadPath(c.Request.URL.Path)
		if err != nil {
			c.Data(200, "text/html; charset=utf-8", []byte(Error(err)))
			return
		}
		if isHtml {
			c.Data(200, "text/html; charset=utf-8", []byte(res))
			return
		}
		c.Data(200, "text/plain; charset=utf-8", []byte(res))
	})
	return r
}
