package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gameMechanicType string

const (
	CORE      gameMechanicType = "core"
	SECONDARY gameMechanicType = "secondary"
	CONTROLS  gameMechanicType = "controls"
)

type GameMechanic struct {
	gorm.Model
	ID                      uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	GameplayDesignID        uuid.UUID        `gorm:"type:uuid;not null" json:"gameplay_design_id"`
	GameplayDesign          GameplayDesign   `json:"gameplay_design"`
	GameMechanicName        string           `json:"game_mechanic_name"`
	GameMechanicDescription string           `json:"game_mechanic_description"`
	GameMechanicType        gameMechanicType `json:"game_mechanic_type"`
}
