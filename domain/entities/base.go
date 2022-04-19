package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey; type:varchar(36)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()

	return
}
