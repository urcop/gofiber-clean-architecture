package models

import "gorm.io/gorm"

type Example struct {
	*gorm.Model
	Row1 string `json:"row1"`
	Row2 string `json:"row2"`
}
