package middlewares

import (
	"fmt"
	"iBox-Server/internal/core"
	"iBox-Server/internal/exception"
	"iBox-Server/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	tokenString, ok := c.Request.Header["Authorization"]
	if !ok {
		fmt.Println("[Auth] 沒有驗證Token")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.NO_TOKEN,
		})
		return
	}

	claims, err := auth.Verify(tokenString[0])
	if err != nil {
		fmt.Println("[Auth] 無效的 Token")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_TOKEN,
		})
		return
	}

	user := core.GetUserByAccount(claims.Account)

	if user == nil || claims.Password != user.Password {
		fmt.Println("[Auth] 帳號密碼錯誤")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INCORRECT_AUTH,
		})
		return
	}

	fmt.Printf("[Auth] 驗證成功 使用者:%s\n", user)
	c.Set("user", user)
	c.Next()
}

func VerifyUser(c *gin.Context) {
	tokenString, ok := c.Request.Header["Authorization"]
	if !ok {
		fmt.Println("[Auth] 沒有驗證Token")
		c.Next()
		return
	}

	claims, err := auth.Verify(tokenString[0])
	if err != nil {
		fmt.Println("[Auth] 無效的 Token")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_TOKEN,
		})
		return
	}

	user := core.GetUserByAccount(claims.Account)

	if user == nil || claims.Password != user.Password {
		fmt.Println("[Auth] 帳號密碼錯誤")
		c.Next()
		return
	}

	fmt.Printf("[Auth] 驗證成功 使用者:%s\n", user)
	c.Set("user", user)
	c.Next()
}
