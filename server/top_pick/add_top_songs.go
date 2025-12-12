package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	AddTopSongsUrlSuffix = "/add_top_songs"
)

func AddTopSongs(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.AddTopSongsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.AddTopSongsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return AddTopSongsFlow(ctx, requestTyped, responseTyped)
}

func AddTopSongsFlow(ctx context.Context, request *pb.AddTopSongsRequest, response *pb.AddTopSongsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	topSongs := request.GetTopSongs() // assumed []*pb.Artist

	if userId == 0 || len(topSongs) == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	// Map protobuf artists to internal struct
	var songList []*module.TopPickSong
	for _, a := range topSongs {
		songList = append(songList, &module.TopPickSong{
			SongName:  a.GetSongName(),
			DiscogsId: a.GetExternalDgId(),
		})
	}

	// db is your *gorm.DB, you may need to inject it
	_, err := module.CreateUserAllTimeTopSongsTx(ctx, userId, songList)
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
