package main

import (
	"iBox-Server/internal/core"
	"iBox-Server/internal/router"
	"iBox-Server/server"
	"iBox-Server/services/auth"
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

var (
	webRouter     *router.WebRouter
	storageRouter *router.StorageRouter
)

func init() {
	// 設定為正式版
	// gin.SetMode(gin.ReleaseMode)

	auth.Init()
	core.InitStorage()
	core.InitUser()

	// 建立路由
	webRouter = router.NewWebRouter(os.Getenv("WEB_FOLDER"), os.Getenv("STORAGE_FOLDER"))
	storageRouter = router.NewStorageRouter(os.Getenv("STORAGE_FOLDER"), os.Getenv("STORAGE_PREFIX"))
}

func main() {
	defer core.SaveStorage()
	defer core.SaveUser()

	// DEBUG:
	// core.CreateUser("Lanstar", "Lanstar", "danny95624268@gmail.com", "aa95624268")
	// core.CreateUser("Lanstar2", "Lanstar2", "danny95624268@gmail.com", "aa95624268")

	// 同步 Web 版本
	watcher, startWatch := webRouter.WatchWeb()
	go startWatch()
	defer watcher.Close()

	// 啟用 Web HTTP 服務
	go server.ListenAndServe(":"+os.Getenv("WEB_PORT"), webRouter)

	// 啟用 File Webdav 服務
	go server.ListenAndServe(":"+os.Getenv("STORAGE_PORT"), storageRouter)

	// WEBDAV
	// http.ListenAndServe(":8080", &webdav.Handler{
	// 	// Prefix:     "",

	// 	FileSystem: webdav.Dir("."),
	// 	LockSystem: webdav.NewMemLS(),
	// })

	//
	var quit chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("關閉伺服器中...")
}
