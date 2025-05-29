package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type projectType string

const (
	SCRIPTED_ORIENTED projectType = "scripted_oriented"
	OBJECT_ORIENTED   projectType = "object_oriented"
)

type Project struct {
	gorm.Model
	ID           uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	ProjectName  string      `json:"project_name"`
	ProjectType  projectType `json:"project_type"`
	CollectiveID uuid.UUID   `gorm:"type:uuid;not null" json:"collective_id"`
	Collective   Collective  `json:"collective"`
}
