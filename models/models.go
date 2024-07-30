package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Original string
	Short    string
}
