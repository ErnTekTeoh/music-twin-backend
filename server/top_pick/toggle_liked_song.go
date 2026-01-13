package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	ToggleLikedSongUrlSuffix = "/toggle_liked_song"
)

func ToggleLikedSong(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.ToggleLikedSongRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.ToggleLikedSongResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return ToggleLikedSongFlow(ctx, requestTyped, responseTyped)
}

func ToggleLikedSongFlow(ctx context.Context, request *pb.ToggleLikedSongRequest, response *pb.ToggleLikedSongResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	topSong := request.GetLikedSong()

	if userId == 0 || topSong == nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	var err error
	if topSong.GetIsLiked() {
		err = module.DeleteLikedSong(ctx, userId, topSong.GetExternalAmId())
	} else {
		_, err = module.AddLikedSong(ctx, userId, topSong)
	}

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Failed to update liked songs, please try again later.")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
