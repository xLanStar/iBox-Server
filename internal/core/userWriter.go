package core

import (
	fastio "github.com/xLanStar/go-fast-io/v2"
)

type userWriter struct {
	*fastio.FileWriter
}

var UserWriter userWriter

func (userWriter) New() userWriter {
	userWriter := userWriter{&fastio.FileWriter{}}
	userWriter.Init()
	return userWriter
}

func (userWriter *userWriter) WriteUser(user *User) {
	userWriter.WriteString(user.Id)
	userWriter.WriteString(user.Name)
	userWriter.WriteString(user.Account)
	userWriter.WriteString(user.Email)
	userWriter.WriteString(user.Password)
}
