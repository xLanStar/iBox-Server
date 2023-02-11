package middlewares

import (
	"fmt"
	"iBox-Server/internal/core"
	"iBox-Server/internal/exception"
	"iBox-Server/services/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 需要 URL 附帶 filderId 參數，否則會
//
// 若有 VerifiedUser 則會驗證使用者權限
//
// 否則僅以公開權限設定做為參考
func VerifyFolderPermission(c *gin.Context) {
	folderId := c.Param("folderid")

	if folderId == "" {
		fmt.Println("folderid = ''")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_FOLDER_ID,
		})
		return
	}

	folder := core.GetFolder(folderId)

	if folder == nil {
		fmt.Println("folder == nil")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_FOLDER_ID,
		})
		return
	}

	user, _ := c.Get("user")

	var folderPermission *permission.FolderPermission
	if user != nil {
		folderPermission = permission.GetFolderPermission(folder, user.(*core.User))
	} else {
		folderPermission = permission.GetFolderPermission(folder, nil)
	}

	c.Set("folder", folder)
	c.Set("permission", folderPermission)
	c.Next()
}

func VerifyFilePermission(c *gin.Context) {
	folderId := c.Param("fileId")

	if folderId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_FOLDER_ID,
		})
		return
	}

	folder := core.GetFolder(folderId)

	if folder == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": exception.INVALID_FOLDER_ID,
		})
		return
	}

	user, _ := c.Get("user")

	var folderPermission *permission.FolderPermission
	if user != nil {
		folderPermission = permission.GetFolderPermission(folder, user.(*core.User))
	} else {
		folderPermission = permission.GetFolderPermission(folder, nil)
	}

	c.Set("permission", folderPermission)
	c.Next()
}
