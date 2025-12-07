package data

import (
	"context"
)

func CreateNewUser(ctx context.Context, newData *User) (*User, error) {
	tx := GetDB().Create(newData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return newData, nil
}

func GetUserByUserId(ctx context.Context, hostId int32) (*User, error) {
	host := &User{}
	tx := GetDB().Where("user_id = ? AND deleted_at IS NULL", hostId).First(host)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return host, nil
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	tx := GetDB().Where("email = ?", email).Where("deleted_at IS NULL").First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func UpdateUser(ctx context.Context, updatedUser *User) error {
	tx := GetDB().Save(updatedUser)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
