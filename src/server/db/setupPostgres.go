package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               uint `gorm:"primaryKey;unique;autoIncrement;"`
	Name             string
	Email            string
	Password         string
	Role             uint        `gorm:"default:1"`
	DirectorateRefer int         `gorm:"default:1"`
	Directorate      Directorate `gorm:"foreignKey:DirectorateRefer"`
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
	Index         int
	UserRefer     int
	ElementsRefer int
	Elements      Elements `gorm:"foreignKey:ElementsRefer"`
	User          User     `gorm:"foreignKey:UserRefer"`
}

type Sensor struct {
	gorm.Model
	Identifier string
	Name       string
	CoordX     float32
	CoordY     float32
}

func SetupPostgres() (*gorm.DB, error) {

	dsn := "host=postgres user=username password=password dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Sensor{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Props{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Elements{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Layout{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Directorate{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	directorate := Directorate{
		ID:          1,
		Directorate: "Default",
	}
	db.Create(&directorate)

	props := Props{
		ID:    1,
		Value: "w-full",
	}
	db.Create(&props)

	elements := [2]Elements{
		{
			ID:         1,
			Name:       "MainChart",
			PropsRefer: 1,
		},
		{
			ID:         2,
			Name:       "HeatMap",
			PropsRefer: 1,
		},
	}
	db.Create(&elements)

	fmt.Println("Inicialização do banco de dados concluída")

	return db, nil
}
