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

	r.HandleFunc(UserEndpointPrefix+UpdateMySocialsSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: UpdateMySocials,
		Request:   &pb.UpdateMySocialsRequest{},
		Response:  &pb.UpdateMySocialsResponse{},
	}))

	r.HandleFunc(UserEndpointPrefix+GetMySocialsSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetMySocials,
		Request:   &pb.GetMySocialsRequest{},
		Response:  &pb.GetMySocialsResponse{},
	}))

	r.HandleFunc(UserEndpointPrefix+GetUserProfileSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetUserProfile,
		Request:   &pb.GetUserProfileRequest{},
		Response:  &pb.GetUserProfileResponse{},
	}))
}
