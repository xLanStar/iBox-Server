package router

import (
	"fmt"
	"iBox-Server/internal/core"
	"iBox-Server/services/auth"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/webdav"
)

type StorageRouter struct {
	webdav.Handler
}

func NewStorageRouter(storageRoot, WebdavPrefix string) *StorageRouter {
	router := &StorageRouter{webdav.Handler{
		Prefix:     WebdavPrefix,
		FileSystem: webdav.Dir(storageRoot),
		LockSystem: webdav.NewMemLS(),
	}}
	return router
}

func (router StorageRouter) Init() {
}

const BaseAuthorizationPrefix = "Basic "

// ServeHTTP conforms to the http.Handler interface.
func (router *StorageRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Method:", req.Method, "URL:", req.URL.Path)

	// CORS
	if origin, ok := req.Header["Origin"]; ok {
		// fmt.Println("Has origin", origin, origin[0])
		w.Header().Set("Access-Control-Allow-Origin", origin[0])
		w.Header().Set("Access-Control-Allow-Headers", "Folder, Authorization, depth") //, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With
		w.Header().Set("Access-Control-Allow-Methods", "PROPFIND")
	}

	// OPTIONS
	if req.Method == "OPTIONS" {
		fmt.Println(req)
		// w.WriteHeader(http.StatusOK)
		return
	}

	// 身分驗證
	var user *core.User

	authorization, hasAuth := req.Header["Authorization"]

	if hasAuth {
		// fmt.Println("[StorageRouter]", "有 Authorization", authorization[0])

		// 有登入驗證
		if strings.HasPrefix(authorization[0], BaseAuthorizationPrefix) {
			// Case 1: 以帳號、密碼登入
			account, password, authOK := req.BasicAuth()
			log.Println("[StorageRouter] 以帳號:", account, " 密碼:", password, "登入")

			if !authOK {
				log.Println("[StorageRouter] 無效的 Authorization")
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user = core.GetUserByAccount(account)

			//
			if user == nil || user.Password != password {
				log.Println("[StorageRouter] 帳號或密碼錯誤")
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else {
			// Case 1: 以Token
			log.Println("[StorageRouter] 以Token 登入")
			authClaim, err := auth.Verify(authorization[0])
			if err != nil {
				log.Println("[StorageRouter] 無效的 Token")
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			user = core.GetUserByAccount(authClaim.Account)
		}
	}

	// DEBUG
	if user != nil {
		fmt.Printf("[StorageRouter] 已登入: %s\n", user)
	}

	// 重導
	if folderId, hasFolderId := req.Header["Folder"]; hasFolderId {
		// Case 1: 指定分享的資料夾
		// TODO: get folder name by folder id
		// TODO: validation share permission
		req.URL.Path = "/" + folderId[0] + req.URL.Path
		log.Println("重導至指定的資料夾:", folderId[0], "最終URL:", req.URL)
	} else if user != nil {
		// Case 2: 使用者自己的資料夾
		req.URL.Path = "/" + user.Name + req.URL.Path
		log.Println("重導至使用者的資料夾:", user.Name, "最終URL:", req.URL)
	} else {
		// Case 3: 沒有指定且沒有登入
		log.Println("Unuthorized: Need Auth")
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	router.Handler.ServeHTTP(w, req)
}
