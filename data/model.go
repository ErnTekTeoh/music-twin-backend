package data

import (
	"music-twin-backend/proto/pb"
	"time"
)

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

	AlternateEmail  *string
	TelegramHandle  *string
	InstagramHandle *string
	WhatsappHandle  *string
}

func (u *User) GetUserID() int {
	if u == nil || u.UserId == nil {
		return 0
	}
	return *u.UserId
}

func (u *User) GetSalt() string {
	if u == nil || u.Salt == nil {
		return ""
	}
	return *u.Salt
}

func (u *User) GetHash() string {
	if u == nil || u.Hash == nil {
		return ""
	}
	return *u.Hash
}

type UserTopPick struct {
	UserTopPickID            *int32  `gorm:"primaryKey;column:user_top_pick_id"`
	UserID                   *int32  `gorm:"index;column:user_id"`
	Type                     *string `gorm:"type:enum('artist','song')"`
	Ranking                  *int32
	AppleMusicArtistImageUrl *string
	AppleMusicArtistName     *string
	AppleMusicExternalID     *string
	AppleMusicSongName       *string
	AppleMusicSongImageUrl   *string
	CreatedAt                *time.Time `gorm:"column:created_at"`
	UpdatedAt                *time.Time `gorm:"column:updated_at"`
	DeletedAt                *time.Time `gorm:"column:deleted_at"`
}

func (u *UserTopPick) IsSong() bool {
	if u == nil || u.Type == nil {
		return false
	}
	return *u.Type == "song"
}

func (u *UserTopPick) IsArtist() bool {
	if u == nil || u.Type == nil {
		return false
	}
	return *u.Type == "artist"
}

type SongSuggestionCard struct {
	SongSuggestionCardID *int32     `gorm:"primaryKey;column:song_suggestion_card_id"`
	SongTitle            *string    `gorm:"column:song_title"`
	ArtistName           *string    `gorm:"column:artist_name"`
	ImageURL             *string    `gorm:"column:image_url"`
	CardHeader           *string    `gorm:"column:card_header"`
	CardSubheader        *string    `gorm:"column:card_subheader"`
	RecommendReason      *string    `gorm:"column:recommend_reason"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

type SongPollCard struct {
	SongPollCardID *int32                `gorm:"primaryKey;column:song_poll_card_id"`
	CardHeader     *string               `gorm:"column:card_header"`
	CardSubheader  *string               `gorm:"column:card_subheader"`
	CardDisclaimer *string               `gorm:"column:card_disclaimer"`
	CreatedAt      *time.Time            `gorm:"column:created_at"`
	UpdatedAt      *time.Time            `gorm:"column:updated_at"`
	Options        []*SongPollCardOption `gorm:"foreignKey:SongPollCardID"`
}

type SongPollCardOption struct {
	SongPollCardOptionID *int32     `gorm:"primaryKey;column:song_poll_card_option_id"`
	SongPollCardID       *int32     `gorm:"index;column:song_poll_card_id"`
	SongTitle            *string    `gorm:"column:song_title"`
	ArtistName           *string    `gorm:"column:artist_name"`
	ImageURL             *string    `gorm:"column:image_url"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

func ToLikedArtists(picks []*UserTopPick) []*pb.LikedArtist {
	var result []*pb.LikedArtist
	for _, utp := range picks {
		if utp == nil || utp.Type == nil || *utp.Type != "artist" {
			continue
		}
		result = append(result, &pb.LikedArtist{
			ArtistName:     utp.AppleMusicArtistName,
			ExternalAmId:   utp.AppleMusicExternalID,
			ArtistImageUrl: utp.AppleMusicArtistImageUrl,
		})
	}
	return result
}

func ToLikedSongs(picks []*UserTopPick) []*pb.LikedSong {
	var result []*pb.LikedSong
	for _, utp := range picks {
		if utp == nil || utp.Type == nil || *utp.Type != "song" {
			continue
		}
		result = append(result, &pb.LikedSong{
			SongName:     utp.AppleMusicSongName,
			ArtistName:   utp.AppleMusicArtistName,
			ExternalAmId: utp.AppleMusicExternalID,
			SongImageUrl: utp.AppleMusicSongImageUrl,
		})
	}
	return result
}
