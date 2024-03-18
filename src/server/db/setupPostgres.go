package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Role 		uint `gorm:"default:1"`
}

type Sensor struct {
	gorm.Model
	Identifier  string
	Name    	string
	CoordX 		float32
	CoordY 		float32
}

func SetupPostgres() *gorm.DB {

	dsn := "host=postgres user=username password=password dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Sensor{})

	return db
}
