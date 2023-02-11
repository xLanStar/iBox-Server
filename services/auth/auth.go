package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/kataras/jwt"
)

var secret []byte

// var cache map[string]*userManager.User = make(map[string]*userManager.User)

func Init() {
	secret = []byte(os.Getenv("SECRET"))
	fmt.Printf("secret: %v\n", secret)
}

func GenerateToken(account, password string) ([]byte, error) {
	userClaims := map[string]interface{}{
		"account":  account,
		"password": password,
		"exp":      time.Now().AddDate(0, 1, 0).Unix(), // NOTE: JWT Token 有效期限一個月
	}

	token, err := jwt.Sign(jwt.HS256, secret, userClaims)

	if err != nil {
		fmt.Printf("[Auth] 簽署 Token 失敗。原因:%s\n", err)
		return nil, err
	}

	return token, err
}

type AuthClaim struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Verify(token string) (*AuthClaim, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, secret, []byte(token))
	if err != nil {
		return nil, err
	}

	var authClaim AuthClaim

	err = verifiedToken.Claims(&authClaim)
	if err != nil {
		return nil, err
	}

	return &authClaim, nil
}
