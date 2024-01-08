// models/user.go
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
