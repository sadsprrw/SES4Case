package models

type Email struct {
	Address string `gorm:"uniqueIndex"`
}
