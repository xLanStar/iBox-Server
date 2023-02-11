package router

import (
	"crypto/md5"
	"fmt"
	"iBox-Server/controllers"
	"iBox-Server/middlewares"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var etag = generateETag()

func generateETag() string {
	data := []byte(time.Now().String())
	return fmt.Sprintf("%x", md5.Sum(data))
}

type WebRouter struct {
	*gin.Engine
	WebRoot     string
	storageRoot string
}

func NewWebRouter(WebRoot, storageRoot string) *WebRouter {
	router := &WebRouter{gin.New(), WebRoot, storageRoot}
	router.Init()
	return router
}

func Test(c *gin.Context) {
	fmt.Println("Test Middleware")
}

func (engine *WebRouter) Init() {
	// Use Gzip
	engine.Engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Cors
	engine.Engine.Use(middlewares.Cors)

	// 快取靜態 Assets 資源
	engine.CachedStatic("/assets", engine.WebRoot+"/assets/", 60*60*24*30, false)
	// engine.Static("/assets", engine.WebRoot+"assets")

	// 快取靜態 Storage 資源
	// engine.CachedStatic("/file", engine.storageRoot, 60*60*24*30, true)

	// API 路由
	apiGroup := engine.Engine.Group("api")
	{
		// Folder
		folderGroup := apiGroup.Group("folder")
		{
			folderGroup.Use(middlewares.VerifyUser, middlewares.VerifyFolderPermission)
			folderGroup.GET("/:folderid", controllers.GetFolder)
			// folderGroup.POST("/:folderid", controllers.CreateFolder)
			// folderGroup.PUT("/:folderid", controllers.EditFolder)
			// folderGroup.DELETE("/:folderid", controllers.DeleteFolder)
		}

		// File
		fileGroup := apiGroup.Group("file")
		{
			fileGroup.Use(middlewares.VerifyUser, middlewares.VerifyFolderPermission)
			// fileGroup.POST("/:folderid", controllers.UploadFile)
			// fileGroup.PUT("/:folderid", controllers.EditFile)
			// fileGroup.DELETE("/:folderid", controllers.DeleteFile)
		}

		// User
		userGroup := apiGroup.Group("user")
		{
			userGroup.POST("/", controllers.Register)
			userGroup.POST("/login", controllers.Login)
			userGroup.POST("/validate", middlewares.RequireAuth, controllers.Validate)
		}
	}

	// Web 路由
	engine.Engine.Use(engine.CachedStaticFileHandler(engine.WebRoot+"/index.html", 60*60*24*30))
}

func (router *WebRouter) WatchWeb() (*fsnotify.Watcher, func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}

	startWatch := func() {
		err = watcher.Add(router.WebRoot)
		if err != nil {
			log.Fatal("Add failed:", err)
		}

		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				etag = generateETag()
				log.Println("Web 資料夾更新，新的版本標記更新為" + etag)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}

	return watcher, startWatch
}

func (router *WebRouter) GetAbsolutePath(relativePath string) string {
	finalPath := path.Join(router.Engine.BasePath(), relativePath)
	if relativePath[len(relativePath)-1] == '/' && finalPath[len(finalPath)-1] != '/' {
		return finalPath + "/"
	}
	return finalPath
}

// 功能相當於有快取功能的
//
//	router.Static("/assets", "./web/assets")
//
// 先將URL去除 prefix 後，取得目標檔案名稱，再去 folder 資料夾找出目標檔案。快取時間維持 duration 秒
//
// 範例：
//
//	router.CachedStatic("/assets", "./web/assets")
func (router *WebRouter) CachedStatic(path, root string, duration int, requiredAuth bool) {
	handler := router.CachedStaticHandler(path, root, duration, requiredAuth)
	pattern := path + "/*f"

	router.Engine.GET(pattern, handler)
	router.Engine.HEAD(pattern, handler)
}

// 快取提供靜態 root 資料夾
//
//	範例：
//
//	// [GET] "/resources/*file"
//	router.Get("/resources", router.CachedStaticHandler("./resources/", 86400))
//	// [ANY] middlewares
//	router.Use(router.CachedStaticHandler("./resources/", 86400))
func (router *WebRouter) CachedStaticHandler(path, root string, duration int, requiredAuth bool) gin.HandlerFunc {
	fileServer := http.StripPrefix(router.GetAbsolutePath(path), http.FileServer(gin.Dir(root, false)))

	cache := "public, max-age=" + strconv.Itoa(duration)
	return func(c *gin.Context) {
		c.Header("Cache-Control", cache)
		c.Header("ETag", etag)

		if match := c.GetHeader("If-None-Match"); strings.Contains(match, etag) {
			c.Status(http.StatusNotModified)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// 快取提供靜態檔案
//
//	範例：
//
//	// [GET] "/"
//	router.Get("/", router.CachedStaticFileHandler("./web/index.html", 86400))
//	// [ANY] middlewares
//	router.Use(router.CachedStaticFileHandler("./web/index.html", 86400))
func (router *WebRouter) CachedStaticFileHandler(filepath string, duration int) gin.HandlerFunc {
	cache := "public, max-age=" + strconv.Itoa(duration)
	return func(c *gin.Context) {
		c.Header("Cache-Control", cache)
		c.Header("ETag", etag)

		if match := c.GetHeader("If-None-Match"); strings.Contains(match, etag) {
			c.Status(http.StatusNotModified)
			return
		}

		c.File(filepath)
	}
}
