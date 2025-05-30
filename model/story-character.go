package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoryCharacter struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
	GameStoryID   uuid.UUID `gorm:"type:uuid;not null" json:"game_story_id"`
	GameStory     GameStory `json:"game_story"`
	CharacterName string    `json:"character_name"`
}

// TODO: all this story stuff should be a different feature
