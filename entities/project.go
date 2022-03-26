package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model `json:"-"`
	Nama string
	Task []Task `gorm:"ForeignKey:Project_ID"`
}
