package controllers

import (
	"iBox-Server/internal/core"
	"iBox-Server/services/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FolderData struct {
	Permission core.Permission `json:"permission"`
	IsOwner    bool            `json:"is_owner"`

	// READ_ONLY
	Folder     *core.Folder   `json:"folder"`
	Files      []string       `json:"files"`
	SubFolders []*core.Folder `json:"subfolders"`
	Parents    []*core.Folder `json:"parents"`

	// ALL
	PublicPermission core.Permission `json:"public_permission,omitempty"`
	SharePermission  core.Permission `json:"share_permission,omitempty"`
	ShareUsers       []string        `json:"share_users,omitempty"`
}

func GetFolder(c *gin.Context) {
	_folder, _ := c.Get("folder")
	folder := _folder.(*core.Folder)

	_folderPermission, _ := c.Get("permission")
	folderPermission := _folderPermission.(*permission.FolderPermission)

	data := &FolderData{
		Permission: folderPermission.Permission,
		IsOwner:    folderPermission.IsOwner,
	}

	if folderPermission.Permission >= core.READ_ONLY {
		data.Folder = folder
		data.Files = folder.Files
		data.SubFolders = folder.SubFolders

		if _, requestParaent := c.GetQuery("parent"); requestParaent {
			parents := make([]*core.Folder, 0, folderPermission.TopDepth)
			if folderPermission.TopDepth != 0 {
				curFolder := folder
				for curFolder != folderPermission.TopReadableFolder {
					curFolder = curFolder.ParentFolder
					parents = append(parents, curFolder)
				}
			}
			data.Parents = parents
		}
	}

	if folderPermission.Permission == core.ALL {
		data.PublicPermission = folder.PublicPermission
		data.ShareUsers = folder.ShareUsers
		data.SharePermission = folder.SharePermission
	}

	c.JSON(http.StatusOK, data)
}
