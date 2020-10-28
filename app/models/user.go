package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// PostReg 注册信息实体
type PostReg struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
