package tests

import (
	"iBox-Server/internal/core"
	"testing"

	"github.com/gin-gonic/gin"
)

var folder = &core.Folder{
	Id:               "awdikjaowidjioAWJdoijaowd",
	Name:             "tests",
	Files:            []string{"a.txt", "awd.texe", "awd.data", "awd.asef", "a.txt", "awd.texe", "awd.data", "awd.asef", "a.txt", "awd.texe", "awd.data", "awd.asef"},
	PublicPermission: 10,
	SharePermission:  45,
	ShareUsers:       []string{"awda", "awdaowfoiao"},
	ParentFolder:     &core.Folder{},
}

type response struct {
	Permission       core.Permission
	IsOwner          bool
	Folder           *core.Folder
	Files            []string
	SubFolders       []*core.Folder
	Parents          []*core.Folder
	PublicPermission core.Permission
	SharePermission  core.Permission
	ShareUsers       []string
}

func BenchmarkResponseMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permission := core.ALL
		fopReadableFolder := folder

		data := gin.H{
			"permission": permission,
			"is_owner":   false,
		}

		if permission >= core.READ_ONLY {
			data["folder"] = folder
			data["subfolders"] = folder.SubFolders

			if true {
				parents := make([]*core.Folder, 0)
				curFolder := folder
				for curFolder != fopReadableFolder {
					curFolder = curFolder.ParentFolder
					parents = append(parents, curFolder)
				}
				data["parents"] = parents
			}
		}

		if permission == 7 {
			data["public_permission"] = folder.PublicPermission
			data["share_users"] = folder.ShareUsers
			data["share_permission"] = folder.SharePermission
		}
	}
}

func BenchmarkResponseStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permission := core.ALL
		fopReadableFolder := folder

		data := &response{}

		if permission >= core.READ_ONLY {
			data.Folder = folder
			data.SubFolders = folder.SubFolders

			if true {
				parents := make([]*core.Folder, 0)
				curFolder := folder
				for curFolder != fopReadableFolder {
					curFolder = curFolder.ParentFolder
					parents = append(parents, curFolder)
				}
				data.Parents = parents
			}
		}

		if permission == 7 {
			data.PublicPermission = folder.PublicPermission
			data.ShareUsers = folder.ShareUsers
			data.SharePermission = folder.SharePermission
		}
	}
}
