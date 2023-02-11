package core

import (
	"fmt"
	"os"

	fastio "github.com/xLanStar/go-fast-io"
)

type folderReader struct {
	*fastio.FileReader
}

var FolderReader folderReader

func (folderReader) New() folderReader {
	folderReader := folderReader{&fastio.FileReader{}}
	folderReader.Init()
	return folderReader
}

func (folderReader *folderReader) ReadFolder() (*Folder, error) {
	id := folderReader.ReadString()
	folderName := folderReader.ReadString()
	share := folderReader.ReadUint8()
	shareUserCount := folderReader.ReadUint8()
	subFolderCount := folderReader.ReadUint32()

	dh, err := os.Open(STORAGE_FOLDER + "/" + id)
	if err != nil {
		fmt.Printf("[FileManager] 資料夾 id:%s name:%s 開啟資料夾失敗\n", id, folderName)
		return nil, err
	}

	fileinfos, err := dh.Readdir(-1)
	if err != nil {
		fmt.Printf("[FileManager] 資料夾 id:%s name:%s 讀取資料夾失敗\n", id, folderName)
		dh.Close()
		return nil, err
	}

	folder := &Folder{
		Id:               id,
		Name:             folderName,
		Files:            make([]string, len(fileinfos)),
		SubFolders:       make([]*Folder, subFolderCount),
		ShareUsers:       make([]string, shareUserCount),
		PublicPermission: Permission(share & 15),
		SharePermission:  Permission((share >> 4) & 15),
		// ShareUsers:       make([]string, 0),
		// IsPublic:         false,
		// PublicPermission: Permission.NO_PERMISSION,
		// SharePermission:  Permission.NO_PERMISSION,
	}

	for i := uint8(0); i != shareUserCount; i++ {
		folder.ShareUsers[i] = folderReader.ReadString()
	}

	for i, fileinfo := range fileinfos {
		folder.Files[i] = fileinfo.Name()
	}

	for i := uint32(0); i != subFolderCount; i++ {
		subFolder, err := folderReader.ReadFolder()
		if err != nil {
			fmt.Printf("[FileManager] 資料夾 id:%s name:%s 讀取子資料夾失敗\n", id, folderName)
			return folder, err
		}
		folder.SubFolders[i] = subFolder
		subFolder.ParentFolder = folder
	}

	storageIdMap[folder.Id] = folder

	return folder, nil
}
