package _struct

import (
	"github.com/jinzhu/gorm"
)

// Specification for an error entry
type Error struct {
	gorm.Model

	EndpointID   string
	ErrorID      uint16
	ErrorMessage string
	Comments     string
}
