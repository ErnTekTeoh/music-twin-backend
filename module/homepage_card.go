package module

import (
	"context"
	"music-twin-backend/data"
)

func GetSongSuggestionCards(ctx context.Context, limit int, offset int) ([]*data.SongSuggestionCard, error) {
	return data.ListSongSuggestionCards(ctx, limit, offset)
}

// GetSongPollCardsWithOptions fetches poll cards and attaches their options
func GetSongPollCardsWithOptions(ctx context.Context, limit int, offset int) ([]*data.SongPollCard, error) {
	cards, err := data.ListSongPollCards(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Attach options for each poll card
	for _, card := range cards {
		if card.SongPollCardID == nil {
			continue
		}
		options, optErr := data.ListSongPollCardOptionsByPollCardID(ctx, *card.SongPollCardID)
		if optErr != nil {
			return nil, optErr
		}
		card.Options = options
	}
	return cards, nil
}
