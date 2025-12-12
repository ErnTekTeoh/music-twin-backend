package server

import (
	"github.com/gorilla/mux"
	"music-twin-backend/server/search"
	"music-twin-backend/server/top_pick"
	"music-twin-backend/server/user"
)

func InitHttpEndpoints(r *mux.Router) {
	user.InitUserEndpoints(r)
	search.InitSearchEndpoints(r)
	top_pick.InitTopPicksEndpoints(r)
}
