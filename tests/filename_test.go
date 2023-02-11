package tests

import (
	"path/filepath"
	"testing"
)

const folderPath = "C:/asd/213/ttsesta.rar"

func BenchmarkA(t *testing.B) {
	var a string
	for i := 0; i != t.N; i++ {
		a = filepath.Base(folderPath)
		_ = a
	}
}

func BenchmarkB(t *testing.B) {
	var a string
	for i := 0; i != t.N; i++ {
		for j := len(folderPath) - 1; j >= 0; j-- {
			if folderPath[j] == '/' {
				a = folderPath[j+1:]
				_ = a
				break
			}
		}
	}
}
