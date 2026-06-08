package models

import "time"

type ProjectTechnology struct {
	ProjectID    int       `json:"project_id" db:"project_id"`
	TechnologyID int       `json:"technology_id" db:"technology_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
