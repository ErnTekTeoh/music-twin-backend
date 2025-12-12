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
