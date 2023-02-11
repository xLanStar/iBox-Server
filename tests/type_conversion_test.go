package tests

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const a = 123456789

func BenchmarkItoa(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(a)
		res = s
	}
	fmt.Println(res)
}
func BenchmarkFastParse(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		builder := strings.Builder{}
		n := a
		for n != 0 {
			builder.WriteRune(rune(48 + (n % 10)))
			n /= 10
		}
		res = builder.String()
	}
	fmt.Println(res)
}
