package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	AddLikedSongsUrlSuffix = "/add_liked_songs"
)

func AddLikedSongs(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.AddLikedSongsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.AddLikedSongsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return AddLikedSongsFlow(ctx, requestTyped, responseTyped)
}

func AddLikedSongsFlow(ctx context.Context, request *pb.AddLikedSongsRequest, response *pb.AddLikedSongsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	topSongs := request.GetLikedSongs() // assumed []*pb.LikedArtis

	if userId == 0 || len(topSongs) == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	// db is your *gorm.DB, you may need to inject it
	_, err := module.CreateUserAllTimeTopSongsTx(ctx, userId, topSongs)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Failed to update top artists, please try again later.")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
