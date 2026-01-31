package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/data"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetUserProfileSuffix = "/get_user_profile"
)

func GetUserProfile(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetUserProfileRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetUserProfileResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetUserProfileFlow(ctx, requestTyped, responseTyped)
}

func GetUserProfileFlow(ctx context.Context, request *pb.GetUserProfileRequest, response *pb.GetUserProfileResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	user, err := module.GetUserDetails(ctx, userId)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error loading user details. Please try again later")
		return
	}

	topSongs, topArtists := module.GetUserTopPicks(ctx, userId)

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.UserDetails = &pb.UserDetails{
		DisplayName:      user.DisplayName,
		ProfileImageUrl:  user.ProfileImageURL,
		UserReferralCode: user.UserReferralCode,
	}

	response.MySocials = &pb.MySocials{
		Telegram:  user.TelegramHandle,
		Whatsapp:  user.WhatsappHandle,
		Instagram: user.InstagramHandle,
		Email:     user.Email,
	}

	response.LikedArtists = data.ToLikedArtists(topArtists)
	response.LikedSongs = data.ToLikedSongs(topSongs)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
