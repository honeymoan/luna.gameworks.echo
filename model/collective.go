package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Collective struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	CollectiveName string    `json:"collective_name"`
}
