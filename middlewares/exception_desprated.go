package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func recoveryHandler(c *gin.Context) {
	if err := recover(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": err,
		})
	}
}

func ExceptionHandler(c *gin.Context) {
	defer recoveryHandler(c)
	c.Next()
}
