package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/data"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetLikedSongsSuffix = "/get_liked_songs"
)

func GetLikedSongs(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetLikedSongsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetLikedSongsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetLikedSongsFlow(ctx, requestTyped, responseTyped)
}

func GetLikedSongsFlow(ctx context.Context, request *pb.GetLikedSongsRequest, response *pb.GetLikedSongsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	topSongs, _ := module.GetUserTopPicks(ctx, userId)
	likedSongs := data.ToLikedSongs(topSongs)

	likedMap := module.GetUserLikedSongsMap(ctx, userId)

	for _, eachSong := range likedSongs {
		eachSong.IsLiked = proto.Bool(false)
		if liked, _ := likedMap[eachSong.GetExternalAmId()]; liked {
			eachSong.IsLiked = proto.Bool(true)
		}
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.LikedSongs = likedSongs
	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
