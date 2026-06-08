package models

import "time"

type PerformanceMetric struct {
	ID           int       `json:"id" db:"id"`
	MetricType   string    `json:"metric_type" db:"metric_type"`
	Value        string    `json:"value" db:"value"`
	Label        string    `json:"label" db:"label"`
	Sublabel     *string   `json:"sublabel,omitempty" db:"sublabel"`
	IconName     string    `json:"icon_name" db:"icon_name"`
	DisplayOrder int       `json:"display_order" db:"display_order"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
