package middlewares

import "github.com/gin-gonic/gin"

func Cors(c *gin.Context) {
	header := c.Writer.Header()
	if origin, ok := c.Request.Header["Origin"]; ok {
		header.Set("Access-Control-Allow-Origin", origin[0])
		header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization") //, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With
	}
	// header.Set("Access-Control-Allow-Credentials", "true")
	// header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

	c.Next()
}
