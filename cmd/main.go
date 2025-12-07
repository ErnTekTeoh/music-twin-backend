package main

import (
	"github.com/gorilla/mux"
	"music-twin-backend/server"
	"net/http"
)

var commitHash string

func main() {
	//
	//if info, ok := debug.ReadBuildInfo(); ok {
	//	for _, setting := range info.Settings {
	//		if setting.Key == "vcs.revision" {
	//			common.LogWithDebug(context.Background(), "Current revision: "+setting.Value)
	//		}
	//	}
	//}
	//
	r := mux.NewRouter()
	//r.HandleFunc("/api/health", HealthCheck)
	//
	//r.HandleFunc("/api/user_profile_image/upload", member.UploadUserProfileImage)
	//
	server.InitHttpEndpoints(r)
	//module.InitPushClient()
	//
	//
	http.ListenAndServe("0.0.0.0:8080", r)
}

// SetCommitHash sets the commit hash during compilation using -ldflags.
// go build -ldflags "-X main.commitHash=`git rev-parse --short HEAD`"
func SetCommitHash(hash string) {
	commitHash = hash
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//_ = json.NewEncoder(w).Encode(&pb.HealthCheckResponse{
	//	Error:        proto.Int32(0),
	//	ErrorMessage: proto.String(commitHash),
	//})
	//return
}
