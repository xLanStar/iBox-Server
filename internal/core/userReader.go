package core

import (
	fastio "github.com/xLanStar/go-fast-io/v2"
)

type userReader struct {
	*fastio.FileReader
}

var UserReader userReader

func (userReader) New() userReader {
	userReader := userReader{&fastio.FileReader{}}
	userReader.Init()
	return userReader
}

func (userReader *userReader) ReadUser() *User {
	id := userReader.ReadString()
	name := userReader.ReadString()
	account := userReader.ReadString()
	email := userReader.ReadString()
	password := userReader.ReadString()
	return &User{
		Id:       id,
		Name:     name,
		Account:  account,
		Email:    email,
		Password: password,
		// RootFolder: GetFolder(folderId),
	}
}
