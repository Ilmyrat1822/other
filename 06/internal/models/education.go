package models

import "time"

type Education struct {
	ID              int       `json:"id" db:"id"`
	InstitutionName string    `json:"institution_name" db:"institution_name"`
	DegreeType      string    `json:"degree_type" db:"degree_type"`
	Status          string    `json:"status" db:"status"`
	StartYear       *int      `json:"start_year,omitempty" db:"start_year"`
	EndYear         *int      `json:"end_year,omitempty" db:"end_year"`
	FieldOfStudy    *string   `json:"field_of_study,omitempty" db:"field_of_study"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
