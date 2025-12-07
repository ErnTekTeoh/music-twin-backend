package search

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/common"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
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

	resp, err := module.SearchArtist(ctx, name)

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error retrieving results. Please try again later")
		return
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Results = ConvertResultsToArtistResp(resp.Results)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}

func ConvertResultsToArtistResp(res []common.SearchResult) []*pb.Artist {
	final := make([]*pb.Artist, 0)
	for _, each := range res {
		final = append(final, &pb.Artist{
			ArtistName:     proto.String(each.Title),
			ArtistImageUrl: proto.String(each.Thumb),
		})
	}
	return final
}
