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

func GetUserTopPicks(ctx context.Context, userId int32) ([]*UserTopPick, error) {
	var datas []*UserTopPick
	tx := GetDB().WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userId).Find(&datas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datas, nil
}
