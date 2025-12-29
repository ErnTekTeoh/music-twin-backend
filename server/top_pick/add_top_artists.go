package top_pick

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	AddTopArtistsUrlSuffix = "/add_top_artists"
)

func AddTopArtists(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.AddTopArtistsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.AddTopArtistsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return AddTopArtistsFlow(ctx, requestTyped, responseTyped)
}

func AddTopArtistsFlow(ctx context.Context, request *pb.AddTopArtistsRequest, response *pb.AddTopArtistsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	topArtists := request.GetTopArtists() // assumed []*pb.Artist

	if userId == 0 || len(topArtists) == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	// Map protobuf artists to internal struct
	var artistList []*module.TopPickArtist
	for _, a := range topArtists {
		artistList = append(artistList, &module.TopPickArtist{
			ArtistName:   a.GetArtistName(),
			DiscogsId:    a.GetExternalDgId(),
			AppleMusicId: a.GetExternalAmId(),
		})
	}

	// db is your *gorm.DB, you may need to inject it
	_, err := module.CreateUserAllTimeTopArtistsTx(ctx, userId, artistList)
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
