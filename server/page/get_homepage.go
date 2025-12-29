package page

import (
	"context"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/data"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
)

const (
	GetHomepageUrlSuffix = "/get_homepage"
)

func GetHomepage(ctx context.Context, request, response interface{}) (errorCode int32) {
	requestTyped, ok := request.(*pb.GetHomepageRequest)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_TYPE)
	}

	responseTyped, ok := response.(*pb.GetHomepageResponse)
	if !ok {
		return int32(pb.Constant_ERROR_CODE_INVALID_RESPONSE_TYPE)
	}

	return GetHomepageFlow(ctx, requestTyped, responseTyped)
}

func GetHomepageFlow(ctx context.Context, request *pb.GetHomepageRequest, response *pb.GetHomepageResponse) (errorCode int32) {
	userId := request.GetRequestMeta().GetUserId()

	if userId == 0 {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM))
		response.ErrorMessage = proto.String("Invalid parameters")
		return int32(pb.Constant_ERROR_CODE_INVALID_REQUEST_PARAM)
	}

	suggestionCards, err := module.GetSongSuggestionCards(ctx, 10, 0)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Failed to load homepage, please try again later.")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	pollCards, err := module.GetSongPollCardsWithOptions(ctx, 10, 0)
	if err != nil {
		response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR))
		response.ErrorMessage = proto.String("Failed to load homepage, please try again later.")
		return int32(pb.Constant_ERROR_CODE_BUSINESS_ERROR)
	}

	homepageCards := make([]*pb.HomepageCard, 0)
	homepageCards = append(homepageCards, ConvertSongSuggestionCardsToHomepageCards(suggestionCards)...)
	homepageCards = append(homepageCards, ConvertSongPollCardsToHomepageCards(pollCards)...)

	response.Error = proto.Int32(int32(pb.Constant_ERROR_CODE_SUCCESS))
	response.ErrorMessage = proto.String("success")
	response.Cards = homepageCards

	return int32(pb.Constant_ERROR_CODE_SUCCESS)
}

func ConvertSongSuggestionCardsToHomepageCards(cards []*data.SongSuggestionCard) []*pb.HomepageCard {
	var out []*pb.HomepageCard
	for _, c := range cards {
		if c == nil {
			continue
		}
		out = append(out, &pb.HomepageCard{
			SongSuggestionCard: &pb.SongSuggestionCard{
				SongTitle:       c.SongTitle,
				ArtistName:      c.ArtistName,
				ImageUrl:        c.ImageURL,
				CardHeader:      c.CardHeader,
				CardSubheader:   c.CardSubheader,
				RecommendReason: c.RecommendReason,
			},
		})
	}
	return out
}

func ConvertSongPollCardsToHomepageCards(cards []*data.SongPollCard) []*pb.HomepageCard {
	var out []*pb.HomepageCard
	for _, c := range cards {
		if c == nil {
			continue
		}
		out = append(out, &pb.HomepageCard{
			SongPollCard: &pb.SongPollCard{
				CardHeader:     c.CardHeader,
				CardSubheader:  c.CardSubheader,
				CardDisclaimer: c.CardDisclaimer,
				Options:        ConvertSongPollCardOptions(c.Options),
			},
		})
	}
	return out
}

func ConvertSongPollCardOptions(opts []*data.SongPollCardOption) []*pb.SongPollCardOption {
	var out []*pb.SongPollCardOption
	for _, o := range opts {
		if o == nil {
			continue
		}
		out = append(out, &pb.SongPollCardOption{
			SongTitle:            o.SongTitle,
			ArtistName:           o.ArtistName,
			ImageUrl:             o.ImageURL,
			SongPollCardOptionId: o.SongPollCardOptionID,
			SongPollCardId:       o.SongPollCardID,
		})
	}
	return out
}
