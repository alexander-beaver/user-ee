package _struct

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Error struct{
	gorm.Model

	EndpointID string
	ErrorID uint16
	ErrorMessage string
	Comments string
	Time time.Time
}