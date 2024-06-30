package models

import "gorm.io/gorm"

type Task struct {
	ID        int            `gorm:"primaryKey"`
	Title     string         `gorm:"type:varchar(100)"`
	Status    string         `gorm:"type:varchar(20)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
