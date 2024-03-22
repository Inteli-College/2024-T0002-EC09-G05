package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               uint `gorm:"primaryKey;unique;autoIncrement;"`
	Name             string
	Email            string
	Password         string
	Role             uint `gorm:"default:1"`
	DirectorateRefer int
	Directorate      Directorate `gorm:"foreignKey:DirectorateRefer"`

	// use LayoutRefer as foreign key
}

type Directorate struct {
	gorm.Model
	ID          int `gorm:"primaryKey;unique;autoIncrement;"`
	Directorate string
}

type Elements struct {
	gorm.Model
	ID         int `gorm:"primaryKey;unique;autoIncrement;"`
	Name       string
	PropsRefer int
	Index      int   `gorm:"unique;autoIncrement;"`
	Props      Props `gorm:"foreignKey:PropsRefer"`
	UserRefer  int
	User       User `gorm:"foreignKey:UserRefer"`
}

type Props struct {
	gorm.Model
	ID    int `gorm:"primaryKey;unique;autoIncrement;"`
	Value string
}

type Sensor struct {
	gorm.Model
	Identifier string
	Name       string
	CoordX     float32
	CoordY     float32
}

func SetupPostgres() *gorm.DB {

	dsn := "host=postgres user=username password=password dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Sensor{})
	db.AutoMigrate(&Props{})
	db.AutoMigrate(&Elements{})
	db.AutoMigrate(&Directorate{})

	return db
}
