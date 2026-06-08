package models

import "time"

type ProjectStatus string

const (
	ProjectStatusLive       ProjectStatus = "LIVE"
	ProjectStatusCompleted  ProjectStatus = "COMPLETED"
	ProjectStatusInProgress ProjectStatus = "IN_PROGRESS"
)

type Project struct {
	ID             int           `json:"id" db:"id"`
	Title          string        `json:"title" db:"title"`
	Description    string        `json:"description" db:"description"`
	Status         ProjectStatus `json:"status" db:"status"`
	Year           int           `json:"year" db:"year"`
	IsFeatured     bool          `json:"is_featured" db:"is_featured"`
	IsConfidential bool          `json:"is_confidential" db:"is_confidential"`
	ProjectURL     *string       `json:"project_url,omitempty" db:"project_url"`
	RepositoryURL  *string       `json:"repository_url,omitempty" db:"repository_url"`
	DisplayOrder   int           `json:"display_order" db:"display_order"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	Technologies   []Technology  `json:"technologies,omitempty" db:"-"`
}
