package page

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetConnectPageUrlSuffix = "/get_connect_page"
)

func GetConnectPage(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetConnectPageRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetConnectPageResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetConnectPageFlow(ctx, requestTyped, responseTyped)
}

func GetConnectPageFlow(ctx context.Context, request *pb.GetConnectPageRequest, response *pb.GetConnectPageResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.CommunityHighlights = module.GetRandomCommunityMockHighlightsFromTopPicks(ctx, userId)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
