package entities

import "github.com/google/uuid"

type WalletType string

const (
	CheckingAccount WalletType = "checking_account"
	SavingAccount   WalletType = "saving_account"
)

type Wallet struct {
	BaseEntity
	Balance      float64    `json:"balance" gorm:"default:0.00"`
	Type         WalletType `json:"type"`
	UserID       uuid.UUID  `gorm:"type:varchar(36)"`
	User         User
	Transactions []Transaction `gorm:"foreignKey:FromWalletID"`
}
