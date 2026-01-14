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

	r.HandleFunc(TopPickEndpointPrefix+ToggleLikedSongUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: ToggleLikedSong,
		Request:   &pb.ToggleLikedSongRequest{},
		Response:  &pb.ToggleLikedSongResponse{},
	}))

	r.HandleFunc(TopPickEndpointPrefix+GetLikedSongsSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetLikedSongs,
		Request:   &pb.GetLikedSongsRequest{},
		Response:  &pb.GetLikedSongsResponse{},
	}))

	r.HandleFunc(TopPickEndpointPrefix+GetLikedArtistsSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetLikedArtists,
		Request:   &pb.GetLikedArtistsRequest{},
		Response:  &pb.GetLikedArtistsResponse{},
	}))

	r.HandleFunc(TopPickEndpointPrefix+ToggleLikedArtistUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: ToggleLikedArtist,
		Request:   &pb.ToggleLikedArtistRequest{},
		Response:  &pb.ToggleLikedArtistResponse{},
	}))
}
