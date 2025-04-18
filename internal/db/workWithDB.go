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

func (r *GormUserRepository) CheckUniqGuid(guid string) bool {
	tmp := User{}
	r.db.Where("guid = ?", guid).First(&tmp)
	return tmp.Guid == ""
}

func (r *GormUserRepository) UpdateReftokenLiveTokenUnicCode(guid, refreshTokenHash, LiveToken, unicCode string) {
	r.db.Model(&User{}).Where("guid = ?", guid).UpdateColumn("refresh_token_hash", refreshTokenHash).
		UpdateColumn("live_token", LiveToken).UpdateColumn("unic_code", unicCode)
}
