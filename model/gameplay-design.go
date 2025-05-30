package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameplayDesign struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	ProjectDesignID  uuid.UUID      `gorm:"type:uuid;not null" json:"project_design_id"`
	ProjectDesign    ProjectDesign  `json:"project_design"`
	CoreGameplayLoop string         `json:"core_gameplay_loop"`
	PlayerObjectives string         `json:"player_objectives"`
	GameProgression  string         `json:"game_progression"`
	GameModes        []GameMode     `json:"game_modes" gorm:"foreignKey:GameplayDesignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GameMechanics    []GameMechanic `json:"game_mechanics" gorm:"foreignKey:GameplayDesignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// TODO: fix all the relationships in the models, including constraint relations
