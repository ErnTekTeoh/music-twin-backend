package page

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/common"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
	"strings"
)

const (
	GetTrendingChartsUrlSuffix = "/get_trending_charts"
)

func GetTrendingCharts(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetTrendingChartsRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetTrendingChartsResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetTrendingChartsFlow(ctx, requestTyped, responseTyped)
}

func GetTrendingChartsFlow(ctx context.Context, request *pb.GetTrendingChartsRequest, response *pb.GetTrendingChartsResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	songs, err := module.GetTop100ChartSongsAndAlbums(ctx)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Error retrieving charts")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	topSongs, topAlbums := ConvertAppleMusicChartResultsToResp(songs)

	likedMap := module.GetUserLikedSongsMap(ctx, userId)

	for _, eachSong := range topSongs {
		eachSong.IsLiked = proto.Bool(false)
		if liked, _ := likedMap[eachSong.GetExternalAmId()]; liked {
			eachSong.IsLiked = proto.Bool(true)
		}
	}
	response.TopSongs = topSongs
	response.TopAlbums = topAlbums

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}

func ConvertAppleMusicChartResultsToResp(res *common.AppleMusicChartAPIResponse) (topSongs []*pb.LikedSong, topAlbums []*pb.TopChartAlbum) {
	finalSongs := make([]*pb.LikedSong, 0)

	finalAlbums := make([]*pb.TopChartAlbum, 0)
	if len(res.Results.Songs) != 0 && res.Results.Songs[0].Data != nil {
		for _, song := range res.Results.Songs[0].Data {
			artworkUrl := song.Attributes.Artwork.Url
			artworkUrl = strings.ReplaceAll(artworkUrl, "{w}", "500")
			artworkUrl = strings.ReplaceAll(artworkUrl, "{h}", "500")
			finalSongs = append(finalSongs, &pb.LikedSong{
				SongName:     proto.String(song.Attributes.Name),
				ArtistName:   proto.String(song.Attributes.ArtistName),
				SongImageUrl: proto.String(artworkUrl),
				ExternalAmId: proto.String(song.ID),
			})
		}
	}

	if len(res.Results.Albums) != 0 && res.Results.Albums[0].Data != nil {
		for _, album := range res.Results.Albums[0].Data {
			artworkUrl := album.Attributes.Artwork.Url
			artworkUrl = strings.ReplaceAll(artworkUrl, "{w}", "500")
			artworkUrl = strings.ReplaceAll(artworkUrl, "{h}", "500")
			finalAlbums = append(finalAlbums, &pb.TopChartAlbum{
				AlbumName:     proto.String(album.Attributes.Name),
				ArtistName:    proto.String(album.Attributes.ArtistName),
				AlbumImageUrl: proto.String(artworkUrl),
				ExternalAmId:  proto.String(album.ID),
			})
		}
	}

	return finalSongs, finalAlbums
}
