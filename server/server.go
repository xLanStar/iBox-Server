package server

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"time"
)

var etag = generateETag()

func generateETag() string {
	data := []byte(time.Now().String())
	return fmt.Sprintf("%x", md5.Sum(data))
}

// 啟用 HTTP Server 服務
func ListenAndServe(addr string, handler http.Handler) {
	log.Printf("開始伺服器服務 %s\n", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil && err != http.ErrServerClosed {
		log.Printf("發生錯誤: %s\n", err)
	}
}

// 啟用 TLS 協定 HTTP Server 服務
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) {
	log.Printf("開始伺服器服務 %s\n", addr)
	err := http.ListenAndServeTLS(addr, certFile, keyFile, handler)
	if err != nil && err != http.ErrServerClosed {
		log.Printf("發生錯誤: %s\n", err)
	}
}
