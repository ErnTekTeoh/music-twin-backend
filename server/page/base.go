package page

import (
	"github.com/gorilla/mux"
	"music-twin-backend/middleware"
	"music-twin-backend/proto/pb"
)

const (
	PageEndpointPrefix = "/api/page"
)

func InitPageEndpoints(r *mux.Router) {

	r.HandleFunc(PageEndpointPrefix+GetHomepageUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetHomepage,
		Request:   &pb.GetHomepageRequest{},
		Response:  &pb.GetHomepageResponse{},
	}))

	r.HandleFunc(PageEndpointPrefix+GetTrendingChartsUrlSuffix, middleware.HttpProcessorWrapper(middleware.HttpProcessor{
		Processor: GetTrendingCharts,
		Request:   &pb.GetTrendingChartsRequest{},
		Response:  &pb.GetTrendingChartsResponse{},
	}))
}
