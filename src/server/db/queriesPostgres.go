package db

import "gorm.io/gorm"

func (u *User) SaveUser(pg *gorm.DB, userData *User) (*User, error) {

	tx := pg.Create(&u)
	if tx.Error != nil {
		return &User{}, tx.Error
	}
	return u, nil
}
