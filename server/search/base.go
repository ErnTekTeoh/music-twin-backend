package search

import (
	"github.com/gorilla/mux"
	"music-twin-backend/middleware"
	"music-twin-backend/proto/pb"
)

const (
	SearchEndpointPrefix = "/api/search"
)

func InitSearchEndpoints(r *mux.Router) {
	// search song based on title + artist + genre
	r.HandleFunc(SearchEndpointPrefix+SearchTrackUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: SearchTrack,
		Request:   &pb.SearchTrackRequest{},
		Response:  &pb.SearchTrackResponse{},
	}))
	// search artist
	r.HandleFunc(SearchEndpointPrefix+SearchArtistUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: SearchArtist,
		Request:   &pb.SearchArtistRequest{},
		Response:  &pb.SearchArtistResponse{},
	}))

}
