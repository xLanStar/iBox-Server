package core

import (
	"fmt"
	"os"
)

var (
	// FolderId: *Folder
	storageIdMap map[string]*Folder = make(map[string]*Folder)
	// *RootFolder: UserId
	mapOwner map[*Folder]*User = make(map[*Folder]*User)

	STORAGE_FOLDER string
	STORAGE_DATA   string
)

// // Private 函數
// func loadFolder(fastio.FileReader) Folder {

// 	fmt.Println("[StorageManager] 讀取資料夾")

// 	// 建立 Folder
// 	node := Folder{
// 		Id:         folderName,
// 		FolderName: "", // NOTE: 由 folders.bin 讀取個資料夾的真正名稱
// 		// FolderPath: folderPath,
// 		SubFolders: make([]Folder, 0),
// 		Files:      make([]File, 0),
// 	}
// 	mapFolder[node.Id] = &node
// 	// folderCounter++

// 	// 讀取 Folder
// 	f, _ := os.Open(folderPath)
// 	fileInfos, _ := f.Readdir(-1)

// 	for _, fileInfo := range fileInfos {
// 		if fileInfo.IsDir() {
// 			// 讀取 SubFolder
// 			node.SubFolders = append(node.SubFolders, loadFolder(filepath.Join(folderPath, fileInfo.Name()), fileInfo.Name()))
// 		} else {
// 			// 讀取 File
// 			file := File{
// 				// Id:       fileCounter,
// 				FileName: fileInfo.Name(),
// 				FilePath: filepath.Join(folderPath, fileInfo.Name()),
// 				Type:     GetType(fileInfo.Name()),
// 			}
// 			node.Files = append(node.Files, file)
// 			// mapFile[Id] = &file
// 			fileCounter++
// 		}
// 	}

// 	return node
// }

// public 函數
func InitStorage() {
	fmt.Println("[StorageManager] 初始化")
	STORAGE_FOLDER = os.Getenv("STORAGE_FOLDER")
	STORAGE_DATA = os.Getenv("STORAGE_DATA")
	// folderPaths = []string{"D:/UltimateExplorer/", "E:/", "./test/", "D:/Share/"}

	if _, err := os.Stat(STORAGE_DATA); err != nil {
		return
	}

	folderReader := FolderReader.New()
	folderReader.OpenFile(STORAGE_DATA, os.O_RDONLY, 0666)
	for folderReader.Available() {
		folder, err := folderReader.ReadFolder()

		if err != nil {
			continue
		}

		fmt.Printf("[StorageManager] 讀取資料夾 %s\n", folder)
	}
	folderReader.Close()
}

func SaveStorage() {
	fmt.Println("[StorageManager] 保存")
	folderWriter := FolderWriter.New()
	folderWriter.OpenFile(STORAGE_DATA, os.O_CREATE, 0666)
	for _, folder := range storageIdMap {
		folderWriter.WriteFolder(folder)
		fmt.Printf("[StorageManager] 保存資料夾 %s\n", folder)
	}
	folderWriter.Close()
}

// func GetFile(fileId uint32) *File {
// 	return mapFile[fileId]
// }

func GenerateFolderId() string {
	id := uuidGenerator.NewV4().String()
	_, found := storageIdMap[id]
	for found {
		id = uuidGenerator.NewV4().String()
		_, found = storageIdMap[id]
	}
	return id
}

func GetFolder(id string) *Folder {
	return storageIdMap[id]
}

func CreateFolder(user *User) *Folder {
	folder := &Folder{
		Id:         GenerateFolderId(),
		Name:       user.Name,
		Files:      make([]string, 0),
		SubFolders: make([]*Folder, 0),
	}

	// // 子資料夾
	// parentFolder, ok := storageIdMap[folderId]
	// if !ok {
	// 	fmt.Printf("[StorageManager] 資料夾 Id:%s Name:%s 找不到對應的父資料夾\n", folderId)
	// 	os.Remove(STORAGE_FOLDER + folderId)
	// 	return nil
	// }
	// parentFolder.SubFolders = append(parentFolder.SubFolders, folder)
	// folder.ParentFolder = parentFolder

	// 根資料夾
	mapOwner[folder] = user

	err := os.Mkdir(STORAGE_FOLDER+"/"+folder.Name, 0666)

	if err != nil {
		fmt.Printf("[StorageManager] 資料夾 Id:%s Name:%s 建立資料夾失敗\n", folder.Id, folder.Name)
		os.Remove(STORAGE_FOLDER + folder.Name)
		return nil
	}

	storageIdMap[folder.Id] = folder

	return folder
}
