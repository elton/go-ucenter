package user

import (
	"testing"

	"crypto/md5"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkBcrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bcrypt.GenerateFromPassword([]byte("12345678"), 10)
	}
}

func BenchmarkMd5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = md5.Sum([]byte("12345678"))
	}
}
