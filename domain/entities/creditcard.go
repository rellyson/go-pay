package entities

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditCard struct {
	BaseEntity
	Limit        float64   `json:"limit" gorm:"default:0.00"`
	CardNumber   string    `json:"card_number" gorm:"not null"`
	ValidUntil   time.Time `json:"valid_until" gorm:"not_null"`
	CVV          int       `json:"cvv" gorm:"not_null"`
	UserID       uuid.UUID `gorm:"type:varchar(36)"`
	User         User
	Transactions []Transaction `gorm:"foreignKey:FromCreditCardID"`
}

func (c *CreditCard) BeforeCreate(tx *gorm.DB) (err error) {

	c.Limit = generateLimit()
	c.ValidUntil = time.Now().AddDate(5, 0, 0)
	c.CardNumber = generateCardNumber()
	c.CVV = generateCVV()

	return
}

func generateLimit() float64 {
	rand.Seed(time.Now().UnixNano())

	rangeLower := 1000.00
	rangeUpper := 99999.00

	generatedLimit := rangeLower + rand.Float64()*(rangeUpper-rangeLower)

	return generatedLimit
}

func generateCardNumber() string {
	rangeLower := 1000
	rangeUpper := 9999

	firstColumnNumber := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	secondColumnNumber := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	thirdColumnNumber := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	fourthColumnNumber := rangeLower + rand.Intn(rangeUpper-rangeLower+1)

	generatedCardNumber := fmt.Sprintf("%v %v %v %v", firstColumnNumber, secondColumnNumber, thirdColumnNumber, fourthColumnNumber)

	return generatedCardNumber
}

func generateCVV() int {
	rangeLower := 100
	rangeUpper := 999

	generatedCvv := rangeLower + rand.Intn(rangeUpper-rangeLower+1)

	return generatedCvv
}
