package db

import (
	"log"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository() *GormUserRepository {
	db, err := Connect()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) AddRecord(user *User) error {
	err := r.db.Create(&User{
		Guid: user.Guid, RefreshTokenHash: user.RefreshTokenHash,
		IpClient: user.IpClient, LiveToken: user.LiveToken, UnicCode: user.UnicCode,
	}).Error
	return err
}

func (r *GormUserRepository) GetRecord(guid string) (User, error) {
	var tmp User
	err := r.db.Where("guid = ?", guid).First(&tmp).Error
	return tmp, err
}
