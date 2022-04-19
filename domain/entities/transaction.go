package entities

import "github.com/google/uuid"

type TransactionType string

const (
	TED    TransactionType = "ted"
	Pix    TransactionType = "pix"
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	BaseEntity
	Value            float64         `json:"value" gorm:"not null"`
	Type             TransactionType `json:"type" gorm:"not null"`
	AdditionalInfo   string          `json:"additional_info"`
	FromWalletID     uuid.UUID       `json:"from_wallet" gorm:"type:varchar(36)"`
	FromCreditCardID uuid.UUID       `json:"from_creditcard" gorm:"type:varchar(36)"`
	To               string          `json:"to" gorm:"not null"`
}
