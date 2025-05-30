package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectFeature struct {
	gorm.Model
	ID                 uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	ProjectDesignID    uuid.UUID     `gorm:"type:uuid;not null" json:"project_design_id"`
	ProjectDesign      ProjectDesign `json:"project_design"`
	FeatureName        string        `json:"feature_name"`
	FeatureDescription string        `json:"feature_description"`
}
