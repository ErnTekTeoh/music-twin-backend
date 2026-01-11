package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	UpdateMySocialsSuffix = "/update_my_socials"
)

func UpdateMySocials(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.UpdateMySocialsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.UpdateMySocialsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return UpdateMySocialsFlow(ctx, requestTyped, responseTyped)
}

func UpdateMySocialsFlow(ctx context.Context, request *pb.UpdateMySocialsRequest, response *pb.UpdateMySocialsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()
	input := request.GetInput()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	var err error
	switch toUpdate := request.GetSocialToUpdate(); toUpdate {
	case "telegram":
		err = module.UpdateTelegramHandle(ctx, userId, input)
	case "whatsapp":
		err = module.UpdateWhatsappHandle(ctx, userId, input)
	case "email":
		err = module.UpdateAlternateEmail(ctx, userId, input)
	case "instagram":
		err = module.UpdateInstagramHandle(ctx, userId, input)
	default:
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error updating socials. Please try again later")
		return
	}
	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)
	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
