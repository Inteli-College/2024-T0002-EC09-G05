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
	LayoutRefer      int
	Layout           Layout `gorm:"foreignKey:LayoutRefer"`

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
	Props      Props `gorm:"foreignKey:PropsRefer"`
}

type Props struct {
	gorm.Model
	ID    int `gorm:"primaryKey;unique;autoIncrement;"`
	Value string
}

type Layout struct {
	gorm.Model
	ID            int `gorm:"primaryKey;"`
	ElementsRefer int
	Elements      Elements `gorm:"foreignKey:ElementsRefer"`
	Index         int
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
	db.AutoMigrate(&Layout{})
	db.AutoMigrate(&Directorate{})

	return db
}
