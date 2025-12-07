package user

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	UserLoginUrlSuffix = "/login"
)

func UserLogin(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.LoginRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.LoginResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return UserLoginFlow(ctx, requestTyped, responseTyped)
}

func UserLoginFlow(ctx context.Context, request *pb.LoginRequest, response *pb.LoginResponse) (errorCode int32) {

	email := request.GetEmail()
	password := request.GetPassword()
	if email == "" || password == "" {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return
	}

	user := module.VerifyUserPassword(ctx, email, password)

	if user == nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Invalid email or password!")
		return
	}

	token := module.GenerateUserToken(ctx, user.GetUserID())

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Token = proto.String(token)

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}
