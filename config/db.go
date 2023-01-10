package config

import (
	"github.com/enescedev/gotodo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dbUrl string) {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Todo{
		Model:       gorm.Model{},
		Title:       "",
		Description: "",
		HasDone:     "",
	})
	DB = db
}

/*
func NewPostgresDB(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("INIT: error_db_initial_conn: ", err)
	}

	return db // db
}

func SetupDBTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalln("INIT: error_db_initial_tables: ", err)
	}
}
*/
