package models

import (
	"time"
)

type Profile struct {
	ID                 int       `json:"id" db:"id"`
	Name               string    `json:"name" db:"name"`
	Username           string    `json:"username" db:"username"`
	Tagline            string    `json:"tagline" db:"tagline"`
	Location           string    `json:"location" db:"location"`
	Bio                string    `json:"bio" db:"bio"`
	Email              string    `json:"email" db:"email"`
	Phone              string    `json:"phone" db:"phone"`
	AvailabilityStatus string    `json:"availability_status" db:"availability_status"`
	GithubURL          *string   `json:"github_url,omitempty" db:"github_url"`
	LinkedinURL        *string   `json:"linkedin_url,omitempty" db:"linkedin_url"`
	WebsiteURL         *string   `json:"website_url,omitempty" db:"website_url"`
	ResumeURL          *string   `json:"resume_url,omitempty" db:"resume_url"`
	ProfileImageURL    string    `json:"profile_image_url" db:"profile_image_url"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}
