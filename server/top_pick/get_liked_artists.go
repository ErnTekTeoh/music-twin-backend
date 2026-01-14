package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/data"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetLikedArtistsSuffix = "/get_liked_artists"
)

func GetLikedArtists(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetLikedArtistsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetLikedArtistsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetLikedArtistsFlow(ctx, requestTyped, responseTyped)
}

func GetLikedArtistsFlow(ctx context.Context, request *pb.GetLikedArtistsRequest, response *pb.GetLikedArtistsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	_, topArtists := module.GetUserTopPicks(ctx, userId)
	likedArtists := data.ToLikedArtists(topArtists)

	for _, each := range likedArtists {
		each.IsLiked = proto.Bool(true)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.LikedArtists = likedArtists
	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
