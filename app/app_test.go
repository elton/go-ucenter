package app

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	db := GetDB()
	if db == nil {
		t.Fatalf("数据库连接错误")
	}
}
