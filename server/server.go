package main

import (
	"gorm.io/gorm"
	"lawencon.com/credential/config"
	"lawencon.com/credential/dao"
)

func main() {
	//init db and inject to dao and service
	g := initDb()
	dao.SetDao(g)

	//init proto
	config.SetProto()
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
