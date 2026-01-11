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
	r.HandleFunc(TopPickEndpointPrefix+AddLikedArtistsUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: AddLikedArtists,
		Request:   &pb.AddLikedArtistsRequest{},
		Response:  &pb.AddLikedArtistsResponse{},
	}))
	// search artist
	r.HandleFunc(TopPickEndpointPrefix+AddLikedSongsUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: AddLikedSongs,
		Request:   &pb.AddLikedSongsRequest{},
		Response:  &pb.AddLikedSongsResponse{},
	}))

}
