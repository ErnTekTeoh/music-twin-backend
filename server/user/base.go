package user

import (
	"github.com/gorilla/mux"
	"music-twin-backend/middleware"
	"music-twin-backend/proto/pb"
)

const (
	UserEndpointPrefix = "/api/user"
)

func InitUserEndpoints(r *mux.Router) {

	// register

	// login

	// add fav artists

	// add fav genres

	// add vibe
	r.HandleFunc(UserEndpointPrefix+UserLoginUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: UserLogin,
		Request:   &pb.LoginRequest{},
		Response:  &pb.LoginResponse{},
	}))

	r.HandleFunc(UserEndpointPrefix+RegisterUserUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: RegisterUser,
		Request:   &pb.RegisterUserRequest{},
		Response:  &pb.RegisterUserResponse{},
	}))

	r.HandleFunc(UserEndpointPrefix+UpdateDisplayNameSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: UpdateDisplayName,
		Request:   &pb.UpdateDisplayNameRequest{},
		Response:  &pb.UpdateDisplayNameResponse{},
	}))

}
