package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	BucketID uint
	Filename string
	Path     string `gorm:"unique"`
}
