package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	RegisterUserUrlSuffix = "/register"
)

func RegisterUser(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.RegisterUserRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.RegisterUserResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return RegisterUserFlow(ctx, requestTyped, responseTyped)
}

func RegisterUserFlow(ctx context.Context, request *pb.RegisterUserRequest, response *pb.RegisterUserResponse) (errorCode int32) {
	email := request.GetEmail()
	password := request.GetPassword()

	if email == "" {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Email is required")
		return
	}

	user, err := module.RegisterUser(ctx, email, password)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String(err.Error())
		return
	}

	if user.GetUserID() == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Register user failed")
		return
	}

	token := module.GenerateUserToken(ctx, user.GetUserID())
	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.IsSuccessful = proto.Bool(true)
	response.Token = proto.String(token)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
