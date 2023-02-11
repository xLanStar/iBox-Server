package tests

import "testing"

var shareUsers = []string{"aawdawd24", "dadannawra24@2q4", "awda244q", "1512521"}
var myAccount = "awda244q"

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, account := range shareUsers {
			if account == myAccount {
				_ = account
				break
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dic := map[string]bool{}
		for _, account := range shareUsers {
			dic[account] = true
		}
	}
}

func BenchmarkMapPointer(b *testing.B) {
	dic := map[*string]bool{}
	for _, account := range shareUsers {
		dic[&account] = true
	}
	for i := 0; i < b.N; i++ {
		account, ok := dic[&myAccount]
		if ok {
			_ = account
		}
	}
}
