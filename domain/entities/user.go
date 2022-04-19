package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseEntity
	FullName  string    `json:"full_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Cpf       string    `json:"cpf" gorm:"unique; not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	Password  string    `gorm:"not null"`
}

func (u *User) IsPasswordValid(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		// encrypts password
		hash, err := hashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hash
	}

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := hashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hash
	}

	return
}

func hashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(bytes), err
}
