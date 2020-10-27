Package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Mobile string
	Password string
}