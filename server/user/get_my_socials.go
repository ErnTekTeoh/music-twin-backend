package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetMySocialsSuffix = "/get_my_socials"
)

func GetMySocials(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetMySocialsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetMySocialsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetMySocialsFlow(ctx, requestTyped, responseTyped)
}

func GetMySocialsFlow(ctx context.Context, request *pb.GetMySocialsRequest, response *pb.GetMySocialsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	user, err := module.GetUserDetails(ctx, userId)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error loading user details. Please try again later")
		return
	}
	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Instagram = user.InstagramHandle
	response.Telegram = user.TelegramHandle
	response.Email = user.AlternateEmail
	response.Whatsapp = user.WhatsappHandle
	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
