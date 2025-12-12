package top_pick

import (
	"github.com/gorilla/mux"
	"music-twin-backend/middleware"
	"music-twin-backend/proto/pb"
)

const (
	TopPickEndpointPrefix = "/api/top_pick"
)

func InitTopPicksEndpoints(r *mux.Router) {
	// search song based on title + artist + genre
	r.HandleFunc(TopPickEndpointPrefix+AddTopArtistsUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: AddTopArtists,
		Request:   &pb.AddTopArtistsRequest{},
		Response:  &pb.AddTopArtistsResponse{},
	}))
	// search artist
	r.HandleFunc(TopPickEndpointPrefix+AddTopSongsUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: AddTopSongs,
		Request:   &pb.AddTopSongsRequest{},
		Response:  &pb.AddTopSongsResponse{},
	}))

}
