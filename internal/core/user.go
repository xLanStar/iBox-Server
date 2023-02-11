package core

import "fmt"

type User struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Account    string  `json:"-"`
	Email      string  `json:"-"`
	Password   string  `json:"-"`
	RootFolder *Folder `json:"-"`
}

func (user *User) String() string {
	return fmt.Sprintf("使用者{ ID=%v 名字=%v }", user.Id, user.Name)
	// return fmt.Sprintf("使用者 {\n\tId:\t\t%v\n\tName:\t\t%v\n\tAccount:\t%v\n\tPassword:\t%v\n\tRootFolder:\t%v\n}\n", user.Id, user.Name, user.Account, user.Password, user.RootFolder)
}
