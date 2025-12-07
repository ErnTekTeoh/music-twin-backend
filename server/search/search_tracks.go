package search

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/common"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	SearchTrackUrlSuffix = "/search_track"
)

func SearchTrack(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.SearchTrackRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.SearchTrackResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return SearchTrackFlow(ctx, requestTyped, responseTyped)
}

func SearchTrackFlow(ctx context.Context, request *pb.SearchTrackRequest, response *pb.SearchTrackResponse) (errorCode int32) {

	title := request.GetTitle()
	artist := request.GetArtist()
	if title == "" && artist == "" {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	resp, err := module.SearchTrack(ctx, title, artist)

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error retrieving results. Please try again later")
		return
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Results = ConvertResultsToResp(resp.Results)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}

func ConvertResultsToResp(res []common.SearchResult) []*pb.Track {
	final := make([]*pb.Track, 0)
	for _, each := range res {
		final = append(final, &pb.Track{
			Title:         proto.String(each.Title),
			TrackImageUrl: proto.String(each.Thumb),
		})
	}
	return final
}
