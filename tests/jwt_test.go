package tests

import (
	"fmt"
	"testing"
	"time"

	jwt2 "github.com/golang-jwt/jwt/v4"
	jwt "github.com/kataras/jwt"
)

const SECRET = "94awd894aw894f651v165xcv8$@)*()%()#%"

var key = []byte(SECRET)

// var key2 = []byte(SECRET + "2")

const account = "a65d1w6aw"
const password = "aw89f4aw6f51aw"

type JWTData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func BenchmarkJWT1(b *testing.B) {
	var tokenS string
	var claims = JWTData{}
	for i := 0; i != b.N; i++ {
		userClaims := map[string]interface{}{
			"account":  account,
			"password": password,
			"exp":      time.Now().AddDate(0, 1, 0).Unix(), // NOTE: JWT Token 有效期限一個月
		}
		token, _ := jwt.Sign(jwt.HS256, key, userClaims)
		tokenS = string(token)
		verifiedToken, err := jwt.Verify(jwt.HS256, key, token)
		if verifiedToken == nil {
			fmt.Println("not valid", err)
			b.Fatal("not valid", err)
		}
		verifiedToken.Claims(&claims)
	}
	fmt.Println(tokenS, claims.Account, claims.Password)
}

func BenchmarkJWT2(b *testing.B) {
	var tokenS string
	var resultAccount, resultPassword string
	for i := 0; i != b.N; i++ {
		token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
			"account":  account,
			"password": password,
			"exp":      time.Now().AddDate(0, 1, 0).Unix(), // NOTE: JWT Token 有效期限一個月
		})
		tokenStr, _ := token.SignedString(key)
		tokenS = tokenStr

		tokenE, _ := jwt2.Parse(tokenS, func(token *jwt2.Token) (interface{}, error) {
			return key, nil
		})

		if claims, ok := tokenE.Claims.(jwt2.MapClaims); ok && tokenE.Valid {
			resultAccount = claims["account"].(string)
			resultPassword = claims["password"].(string)
			_ = claims
		} else {
			_ = ok
		}
	}
	fmt.Println(tokenS, resultAccount, resultPassword)
}
