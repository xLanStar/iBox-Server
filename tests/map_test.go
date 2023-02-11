package tests

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	a := map[string]bool{"a": true}
	_, ok := a["a"]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
	_, ok = a["b"]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
