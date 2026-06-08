package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActivityType represents the type of activity
type ActivityType string

const (
	ActivityTypeLearning    ActivityType = "LEARNING"
	ActivityTypeProject     ActivityType = "PROJECT"
	ActivityTypeAchievement ActivityType = "ACHIEVEMENT"
	ActivityTypeMilestone   ActivityType = "MILESTONE"
)

// Activity represents a portfolio activity or timeline entry
type Activity struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type         ActivityType       `json:"type" bson:"type"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Icon         string             `json:"icon" bson:"icon"`
	Tags         []string           `json:"tags" bson:"tags"`
	Metadata     ActivityMetadata   `json:"metadata" bson:"metadata"`
	Timestamp    time.Time          `json:"timestamp" bson:"timestamp"`
	DisplayOrder int                `json:"display_order" bson:"display_order"`
}

// ActivityMetadata contains flexible metadata for different activity types
type ActivityMetadata struct {
	// For PROJECT activities
	ClientSatisfaction *int `json:"client_satisfaction,omitempty" bson:"client_satisfaction,omitempty"`
	ProjectID          *int `json:"project_id,omitempty" bson:"project_id,omitempty"`

	// For ACHIEVEMENT activities
	Category *string `json:"category,omitempty" bson:"category,omitempty"`
	BadgeURL *string `json:"badge_url,omitempty" bson:"badge_url,omitempty"`

	// For LEARNING activities
	Technology      *string `json:"technology,omitempty" bson:"technology,omitempty"`
	ProgressPercent *int    `json:"progress_percent,omitempty" bson:"progress_percent,omitempty"`

	// For MILESTONE activities
	MetricValue *string `json:"metric_value,omitempty" bson:"metric_value,omitempty"`
	MetricType  *string `json:"metric_type,omitempty" bson:"metric_type,omitempty"`

	// Common fields
	URL  *string `json:"url,omitempty" bson:"url,omitempty"`
	Year *int    `json:"year,omitempty" bson:"year,omitempty"`
}

// Setting represents dynamic application settings
type Setting struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Key       string             `json:"key" bson:"key"`
	Value     interface{}        `json:"value" bson:"value"`
	Category  string             `json:"category" bson:"category"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// CreateActivityInput is used for creating new activities
type CreateActivityInput struct {
	Type         ActivityType     `json:"type" binding:"required"`
	Title        string           `json:"title" binding:"required"`
	Description  string           `json:"description" binding:"required"`
	Icon         string           `json:"icon" binding:"required"`
	Tags         []string         `json:"tags"`
	Metadata     ActivityMetadata `json:"metadata"`
	DisplayOrder int              `json:"display_order"`
}

// UpdateActivityInput is used for updating activities
type UpdateActivityInput struct {
	Title        *string           `json:"title,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Icon         *string           `json:"icon,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	Metadata     *ActivityMetadata `json:"metadata,omitempty"`
	DisplayOrder *int              `json:"display_order,omitempty"`
}
