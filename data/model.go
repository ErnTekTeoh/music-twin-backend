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

type UserTopPick struct {
	UserTopPickID     *int32  `gorm:"primaryKey;column:user_top_pick_id"`
	UserID            *int32  `gorm:"index;column:user_id"`
	Type              *string `gorm:"type:enum('artist','song')"`
	PeriodType        *string `gorm:"type:enum('all_time','week','month');column:period_type"`
	Year              *int32
	Week              *int32
	Month             *int32
	Ranking           *int32
	ItemID            *int32     `gorm:"column:item_id"`
	DiscogsItemName   *string    `gorm:"column:discogs_item_name"`
	DiscogsExternalID *int32     `gorm:"column:discogs_external_id"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
}
