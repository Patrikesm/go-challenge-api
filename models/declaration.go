package models

import "gorm.io/gorm"

type Declaration struct {
	gorm.Model
	Author    string `json:"author"`
	Image     string `json:"image"`
	Testimony string `json:"testimony"`
}
