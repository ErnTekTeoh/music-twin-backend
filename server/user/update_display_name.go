package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	UpdateDisplayNameSuffix = "/update_display_name"
)

func UpdateDisplayName(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.UpdateDisplayNameRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.UpdateDisplayNameResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return UpdateDisplayNameFlow(ctx, requestTyped, responseTyped)
}

func UpdateDisplayNameFlow(ctx context.Context, request *pb.UpdateDisplayNameRequest, response *pb.UpdateDisplayNameResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	name := request.GetDisplayName()
	if userId == 0 || name == "" {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	err := module.UpdateDisplayName(ctx, userId, name)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error updating name. Please try again later")
		return
	}
	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)
	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
