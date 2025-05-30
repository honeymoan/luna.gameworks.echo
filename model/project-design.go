package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectDesign struct {
	gorm.Model
	ID             uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	ProjectID      uuid.UUID        `gorm:"type:uuid;not null" json:"project_id"`
	Project        Project          `json:"project"`
	Introduction   string           `json:"introduction"`
	ElevatorPitch  string           `json:"elevator_pitch"`
	Genre          string           `json:"genre"`
	Platform       string           `json:"platform"`
	TargetAudience string           `json:"target_audience"`
	Monetization   string           `json:"monetization"`
	FeatureList    []ProjectFeature `json:"feature_list"`
	ArtStyle       string           `json:"art_style"`
	OverallMood    string           `json:"overall_mood"`
	SpecificIdeas  string           `json:"specific_ideas"`
}

// TODO: genre, platform should be a selectable list of genres, not just a string
