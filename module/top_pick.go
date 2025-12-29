package module

import (
	"context"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"music-twin-backend/data"
	"time"
)

type TopPickArtist struct {
	ArtistName   string
	DiscogsId    int32
	AppleMusicId string
}

type TopPickSong struct {
	SongName     string
	DiscogsId    int32
	AppleMusicId string
}

func CreateUserAllTimeTopArtistsTx(ctx context.Context, userId int32, artists []*TopPickArtist) ([]*data.UserTopPick, error) {
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)

	var picks []*data.UserTopPick

	err := data.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, artist := range artists {
			if i >= 3 {
				break
			}
			topPick := &data.UserTopPick{
				UserID:               proto.Int32(userId),
				Type:                 proto.String("artist"),
				DiscogsItemName:      proto.String(artist.ArtistName),
				DiscogsExternalID:    proto.Int32(artist.DiscogsId),
				AppleMusicItemName:   proto.String(artist.ArtistName),
				AppleMusicExternalID: proto.String(artist.AppleMusicId),
				CreatedAt:            &now,
				UpdatedAt:            &now,
			}
			if err := tx.Create(topPick).Error; err != nil {
				return err // rollback transaction
			}
			picks = append(picks, topPick)
		}
		return nil // commit transaction
	})

	if err != nil {
		return nil, err
	}
	return picks, nil
}

func CreateUserAllTimeTopSongsTx(ctx context.Context, userId int32, songs []*TopPickSong) ([]*data.UserTopPick, error) {
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)

	var picks []*data.UserTopPick

	err := data.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, song := range songs {
			if i >= 3 {
				break
			}
			topPick := &data.UserTopPick{
				UserID:               proto.Int32(userId),
				Type:                 proto.String("song"),
				DiscogsItemName:      proto.String(song.SongName),
				DiscogsExternalID:    proto.Int32(song.DiscogsId),
				AppleMusicItemName:   proto.String(song.SongName),
				AppleMusicExternalID: proto.String(song.AppleMusicId),
				CreatedAt:            &now,
				UpdatedAt:            &now,
			}
			if err := tx.Create(topPick).Error; err != nil {
				return err // rollback transaction
			}
			picks = append(picks, topPick)
		}
		return nil // commit transaction
	})

	if err != nil {
		return nil, err
	}
	return picks, nil
}

//
//func CreateUserAllTimeTopArtists(ctx context.Context, userId int32, artists []*TopPickArtist) (*data.UserTopPick, error) {
//	//init the loc
//	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
//	//set timezone,
//	timeNow := time.Now().In(loc)
//
//	newData, createErr := data.CreateNewUserTopPick(ctx, &data.UserTopPick{
//		UserID:            proto.Int32(userId),
//		Type:              proto.String("artist"),
//		DiscogsItemName:   proto.String(discogsItemName),
//		DiscogsExternalID: proto.Int32(discogsExternalId),
//		CreatedAt:         &timeNow,
//		UpdatedAt:         &timeNow,
//	})
//	if createErr != nil {
//		return nil, createErr
//	}
//	return newData, createErr
//}
//
//func CreateUserAllTimeTopSongs(ctx context.Context, userId int32, discogsItemName string, discogsExternalId int32) (*data.UserTopPick, error) {
//	//init the loc
//	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
//	//set timezone,
//	timeNow := time.Now().In(loc)
//
//	newData, createErr := data.CreateNewUserTopPick(ctx, &data.UserTopPick{
//		UserID:            proto.Int32(userId),
//		Type:              proto.String("artist"),
//		DiscogsItemName:   proto.String(discogsItemName),
//		DiscogsExternalID: proto.Int32(discogsExternalId),
//		CreatedAt:         &timeNow,
//		UpdatedAt:         &timeNow,
//	})
//	if createErr != nil {
//		return nil, createErr
//	}
//	return newData, createErr
//}
