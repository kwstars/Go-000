package module

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Passwrod string
}
