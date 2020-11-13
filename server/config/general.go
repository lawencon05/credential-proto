package config

import (
	"log"
)

// CatchErrorGeneral for catch without param
func CatchErrorGeneral() {
	if err := recover(); err != nil {
		log.Println("Error =>", err)
	}
}
