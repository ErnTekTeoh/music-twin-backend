package search

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/common"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
	"strings"
)

const (
	SearchArtistUrlSuffix = "/search_artist"
)

func SearchArtist(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.SearchArtistRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.SearchArtistResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return SearchArtistFlow(ctx, requestTyped, responseTyped)
}

func SearchArtistFlow(ctx context.Context, request *pb.SearchArtistRequest, response *pb.SearchArtistResponse) (errorCode int32) {

	name := request.GetName()
	if name == "" {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	resp, err := module.SearchAppleMusicArtist(ctx, name)

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error retrieving results. Please try again later")
		return
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Results = ConvertAppleMusicArtistResultsToResp(resp)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}

func ConvertAppleMusicArtistResultsToResp(res *common.AppleMusicSearchResponse) []*pb.LikedArtist {
	final := make([]*pb.LikedArtist, 0)

	for _, artist := range res.Results.Artists.Data {
		artworkUrl := artist.Attributes.Artwork.Url
		artworkUrl = strings.ReplaceAll(artworkUrl, "{w}", "500")
		artworkUrl = strings.ReplaceAll(artworkUrl, "{h}", "500")
		final = append(final, &pb.LikedArtist{
			ArtistName:     proto.String(artist.Attributes.Name),
			ArtistImageUrl: proto.String(artworkUrl),
			ExternalAmId:   proto.String(artist.ID),
		})
	}

	return final
}
