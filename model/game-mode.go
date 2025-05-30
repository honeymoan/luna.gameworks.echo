package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameMode struct {
	gorm.Model
	ID                  uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	GameplayDesignID    uuid.UUID      `gorm:"type:uuid;not null" json:"gameplay_design_id"`
	GameplayDesign      GameplayDesign `json:"gameplay_design"`
	GameModeName        string         `json:"game_mode_name"`
	GameModeDescription string         `json:"game_mode_description"`
}
