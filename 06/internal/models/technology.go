package models

import "time"

type Technology struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Category         string    `json:"category" db:"category"`
	ProficiencyLevel string    `json:"proficiency_level" db:"proficiency_level"`
	IsCoreTechnology bool      `json:"is_core_technology" db:"is_core_technology"`
	LearningStatus   *string   `json:"learning_status,omitempty" db:"learning_status"`
	IconURL          *string   `json:"icon_url,omitempty" db:"icon_url"`
	DisplayOrder     int       `json:"display_order" db:"display_order"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
