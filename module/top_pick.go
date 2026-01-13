package module

import (
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"music-twin-backend/data"
	"music-twin-backend/proto/pb"
	"time"
)

func AddLikedSong(ctx context.Context, userId int32, song *pb.LikedSong) (*data.UserTopPick, error) {
	likedSongs, err := AddLikedSongs(ctx, userId, []*pb.LikedSong{song})
	if err != nil {
		return nil, err
	}

	if len(likedSongs) == 0 {
		return nil, errors.New("Error liking song")
	}
	return likedSongs[0], nil
}

func DeleteLikedSong(ctx context.Context, userId int32, externalAppleId string) error {
	likedSongs, err := data.GetLikedSong(ctx, userId, externalAppleId)
	if err != nil {
		return err
	}

	if likedSongs == nil || len(likedSongs) == 0 {
		return errors.New("Liked song not found")
	}

	timeNow := time.Now()
	for _, each := range likedSongs {
		each.DeletedAt = &timeNow
		err = data.UpdateTopPick(ctx, each)
	}

	if err != nil {
		return err
	}
	return nil
}

func CreateUserAllTimeTopArtistsTx(ctx context.Context, userId int32, artists []*pb.LikedArtist) ([]*data.UserTopPick, error) {
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)

	var picks []*data.UserTopPick

	err := data.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, artist := range artists {
			if i >= 3 {
				break
			}
			topPick := &data.UserTopPick{
				UserID:                   proto.Int32(userId),
				Type:                     proto.String("artist"),
				AppleMusicArtistImageUrl: artist.ArtistImageUrl,
				AppleMusicArtistName:     artist.ArtistName,
				AppleMusicExternalID:     artist.ExternalAmId,
				CreatedAt:                &now,
				UpdatedAt:                &now,
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

func AddLikedSongs(ctx context.Context, userId int32, songs []*pb.LikedSong) ([]*data.UserTopPick, error) {
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)

	var picks []*data.UserTopPick

	err := data.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, song := range songs {
			if i >= 3 {
				break
			}
			topPick := &data.UserTopPick{
				UserID:                 proto.Int32(userId),
				Type:                   proto.String("song"),
				AppleMusicArtistName:   song.ArtistName,
				AppleMusicExternalID:   song.ExternalAmId,
				AppleMusicSongName:     song.SongName,
				AppleMusicSongImageUrl: song.SongImageUrl,
				CreatedAt:              &now,
				UpdatedAt:              &now,
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

func GetUserTopPicks(ctx context.Context, userId int32) (topSongs []*data.UserTopPick, topArtists []*data.UserTopPick) {
	allData, err := data.GetUserTopPicks(ctx, userId)
	artists := make([]*data.UserTopPick, 0)
	songs := make([]*data.UserTopPick, 0)
	if err != nil {
		return songs, artists
	}
	for _, each := range allData {
		if each.IsSong() {
			songs = append(songs, each)
		} else if each.IsArtist() {
			artists = append(artists, each)
		}
	}
	return songs, artists
}

func GetUserLikedSongsMap(ctx context.Context, userId int32) (likedMap map[string]bool) {
	userLikedSongs := GetUserLikedSongs(ctx, userId)
	likedMap = make(map[string]bool)
	for _, eachLikedSong := range userLikedSongs {
		likedMap[eachLikedSong.GetAppleMusicExternalID()] = true
	}

	return likedMap
}

func GetUserLikedSongs(ctx context.Context, userId int32) (likedSongs []*data.UserTopPick) {
	allData, err := data.GetUserTopPicks(ctx, userId)
	songs := make([]*data.UserTopPick, 0)
	if err != nil {
		return songs
	}
	for _, each := range allData {
		if each.IsSong() {
			songs = append(songs, each)
		}
	}
	return songs
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
