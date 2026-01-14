package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	ToggleLikedArtistUrlSuffix = "/toggle_liked_artist"
)

func ToggleLikedArtist(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.ToggleLikedArtistRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.ToggleLikedArtistResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return ToggleLikedArtistFlow(ctx, requestTyped, responseTyped)
}

func ToggleLikedArtistFlow(ctx context.Context, request *pb.ToggleLikedArtistRequest, response *pb.ToggleLikedArtistResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	topArtist := request.GetLikedArtist()

	if userId == 0 || topArtist == nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	var err error
	if topArtist.GetIsLiked() {
		err = module.DeleteLikedArtist(ctx, userId, topArtist.GetExternalAmId())
	} else {
		_, err = module.AddLikedArtist(ctx, userId, topArtist)
	}

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Failed to update liked artist, please try again later.")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
