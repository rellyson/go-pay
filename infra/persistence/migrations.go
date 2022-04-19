package persistence

import (
	"github.com/rellyson/go-pay/domain/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Wallet{},
		&entities.CreditCard{},
		&entities.Transaction{},
	)
}
