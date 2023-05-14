package model

import (
	"time"
)

type RegisterCode struct {
	Code      string `gorm:"column:code;type:string;size:64;primaryKey"`
	UseBefore time.Time
	Frequency int    `gorm:"column:frequency;type:int;size:32"`
	Comment   string `gorm:"type:string;size:255;"`
}
