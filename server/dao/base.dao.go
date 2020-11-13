package dao

import (
	"fmt"

	"gorm.io/gorm"
)

var g *gorm.DB

//SetDao for gorm njection
func SetDao(gDB *gorm.DB) {
	g = gDB
}

// CatchError for catch with param error
func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}
