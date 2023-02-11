package core

import "fmt"

type Folder struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Files            []string   `json:"-"`
	SubFolders       []*Folder  `json:"-"`
	PublicPermission Permission `json:"-"`
	SharePermission  Permission `json:"-"`
	ShareUsers       []string   `json:"-"`
	ParentFolder     *Folder    `json:"-"`
}

func (folder *Folder) String() string {
	return fmt.Sprintf("資料夾{ ID=%v 名字=%v }", folder.Id, folder.Name)
	// return fmt.Sprintf("使用者 {\n\tId:\t\t%v\n\tName:\t\t%v\n\tAccount:\t%v\n\tPassword:\t%v\n\tRootFolder:\t%v\n}\n", user.Id, user.Name, user.Account, user.Password, user.RootFolder)
}
