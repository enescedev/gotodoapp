package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title" example:"sport"`
	Description string `json:"description" example:"play football"`
	HasDone     string `json:"has_done" example:"done"`
}

// db.AutoMigrate(&models.Todo{})
