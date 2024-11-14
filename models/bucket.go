package models

import "gorm.io/gorm"

type Bucket struct {
	gorm.Model
	Name      string     `json:"bucket_name" gorm:"unique"`
	Path      string     `json:"bucket_path" gorm:"unique"`
	Documents []Document `json:"documents"`
}
