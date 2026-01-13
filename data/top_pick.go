package data

import (
	"context"
)

func CreateNewUserTopPick(ctx context.Context, newData *UserTopPick) (*UserTopPick, error) {
	tx := GetDB().Create(newData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return newData, nil
}

func GetLikedSong(ctx context.Context, userId int32, externalAppleId string) ([]*UserTopPick, error) {
	var datas []*UserTopPick
	tx := GetDB().WithContext(ctx).Where("user_id = ? AND apple_music_external_id = ? AND type = 'song' AND deleted_at IS NULL", userId, externalAppleId).Find(&datas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datas, nil
}

func GetUserTopPicks(ctx context.Context, userId int32) ([]*UserTopPick, error) {
	var datas []*UserTopPick
	tx := GetDB().WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userId).Find(&datas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datas, nil
}

func UpdateTopPick(ctx context.Context, updatedTopPick *UserTopPick) error {
	tx := GetDB().Save(updatedTopPick)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
