package data

import "time"

type User struct {
	UserId               *int    `gorm:"primaryKey;autoIncrement"`
	DisplayName          *string `gorm:"type:varchar(100)"`
	Email                *string `gorm:"type:varchar(100);index"`
	Salt                 *string `gorm:"type:varchar(16)"`
	Hash                 *string `gorm:"type:varchar(64)"`
	ProfileImageURL      *string `gorm:"type:varchar(1024)"`
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	DeletedAt            *time.Time
	Bio                  *string `gorm:"type:varchar(1024)"`
	Gender               *int
	Location             *string `gorm:"type:varchar(128)"`
	UserReferralCode     *string `gorm:"type:varchar(10);unique"`
	JoiningReferralCode  *string `gorm:"type:varchar(10)"`
	FavouriteArtist1ID   *int
	FavouriteArtist1Name *string `gorm:"type:varchar(128)"`
	FavouriteArtist2ID   *int
	FavouriteArtist2Name *string `gorm:"type:varchar(128)"`
	FavouriteArtist3ID   *int
	FavouriteArtist3Name *string `gorm:"type:varchar(128)"`
	FavouriteGenreName1  *string `gorm:"type:varchar(128)"`
	FavouriteGenreName2  *string `gorm:"type:varchar(128)"`
	FavouriteGenreName3  *string `gorm:"type:varchar(128)"`
}

func (u *User) GetUserID() int {
	return *u.UserId
}

func (u *User) GetSalt() string {
	return *u.Salt
}

func (u *User) GetHash() string {
	return *u.Hash
}
