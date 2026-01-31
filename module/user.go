package module

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"music-twin-backend/common"
	"music-twin-backend/data"
	"time"
)

func RegisterUser(ctx context.Context, email, password string) (user *data.User, err error) {
	//init the loc
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	//set timezone,
	timeNow := time.Now().In(loc)

	existingUser, _ := data.GetUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("Email already in use!")
	}

	// salt and hash logic
	salt := common.GenerateSalt()
	finalHash := HashUserPassword(password, salt)
	newUser, createErr := data.CreateNewUser(ctx, &data.User{
		Email:            proto.String(email),
		Salt:             proto.String(salt),
		Hash:             proto.String(finalHash),
		CreatedAt:        &timeNow,
		UpdatedAt:        &timeNow,
		UserReferralCode: proto.String(GenerateUserReferralCode()),
	})
	if createErr != nil {
		return nil, createErr
	}
	return newUser, createErr
}

func HashUserPassword(password, salt string) string {
	preHash := password + salt
	postHash := sha256.Sum256([]byte(preHash))
	finalHash := hex.EncodeToString(postHash[:])
	return finalHash
}

func VerifyUserPassword(ctx context.Context, email string, password string) *data.User {
	user, err := data.GetUserByEmail(ctx, email)
	if err != nil {
		return nil
	}

	if user == nil {
		return nil
	}

	salt := user.GetSalt()
	hashedPassword := HashUserPassword(password, salt)
	if hashedPassword != user.GetHash() {
		return nil
	}
	return user
}

func UpdateDisplayName(ctx context.Context, userId int32, name string) error {
	user, err := data.GetUserByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.DisplayName = proto.String(name)
	t := common.GetTimeNow()
	user.UpdatedAt = &t

	err = data.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAlternateEmail(ctx context.Context, userId int32, email string) error {
	user, err := data.GetUserByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.AlternateEmail = proto.String(email)
	t := common.GetTimeNow()
	user.UpdatedAt = &t

	err = data.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTelegramHandle(ctx context.Context, userId int32, handle string) error {
	user, err := data.GetUserByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.TelegramHandle = proto.String(handle)
	t := common.GetTimeNow()
	user.UpdatedAt = &t

	err = data.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateInstagramHandle(ctx context.Context, userId int32, handle string) error {
	user, err := data.GetUserByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.InstagramHandle = proto.String(handle)
	t := common.GetTimeNow()
	user.UpdatedAt = &t

	err = data.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWhatsappHandle(ctx context.Context, userId int32, handle string) error {
	user, err := data.GetUserByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.WhatsappHandle = proto.String(handle)
	t := common.GetTimeNow()
	user.UpdatedAt = &t

	err = data.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserDetails(ctx context.Context, userId int32) (*data.User, error) {
	return data.GetUserByUserId(ctx, userId)
}

func GenerateUserReferralCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
