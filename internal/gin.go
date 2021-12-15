package internal

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CreateEngine(dir string, isDisableCors bool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	if !isDisableCors {
		r.Use(CORSMiddleware())
	}
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
