package permission

import (
	"iBox-Server/internal/core"
)

type FolderPermission struct {
	TopReadableFolder *core.Folder
	TopDepth          uint16
	IsOwner           bool
	Permission        core.Permission
	Folder            *core.Folder
}

// 必須有 params /:folderId
func GetFolderPermission(folder *core.Folder, user *core.User) *FolderPermission {

	// 檢查權限
	topReadableFolder := folder
	topDepth := uint16(0)
	permission := core.NO_PERMISSION

	// 使用者相關設定
	// _user, isUser := c.Get("user")
	isOwner := false
	if user != nil {
		curFolder := folder
		depth := uint16(0)
		for {
			for _, shareUser := range curFolder.ShareUsers {
				if shareUser == user.Account {
					if curFolder.SharePermission > permission {
						permission = curFolder.SharePermission
					}
					if curFolder.SharePermission >= core.READ_ONLY {
						topReadableFolder = curFolder
						topDepth = depth
					}
					break
				}
			}
			if curFolder.ParentFolder == nil || permission == core.ALL {
				break
			}
			depth++
			curFolder = curFolder.ParentFolder
		}

		if curFolder == user.RootFolder {
			// 擁有者
			permission = core.ALL
			isOwner = true
		}
	}

	if permission != core.ALL {
		// 公開設定
		curFolder := folder
		depth := uint16(0)
		for {
			if curFolder.PublicPermission >= core.READ_ONLY && depth > topDepth {
				topReadableFolder = curFolder
				topDepth = depth
			}
			if curFolder.PublicPermission > permission {
				permission = curFolder.PublicPermission
				if permission == core.ALL {
					break
				}
			}
			if curFolder.ParentFolder == nil {
				break
			}
			depth++
			curFolder = curFolder.ParentFolder
		}
	}

	// c.Set("permission", FolderPermission{
	// 	TopReadableFolder: topReadableFolder,
	// 	TopDepth:          topDepth,
	// 	IsOwner:           isOwner,
	// 	Permission:        permission,
	// 	Folder:            folder,
	// })
	return &FolderPermission{
		TopReadableFolder: topReadableFolder,
		TopDepth:          topDepth,
		IsOwner:           isOwner,
		Permission:        permission,
	}
}
