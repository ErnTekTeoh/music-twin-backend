package data

import "context"

func CreateSongSuggestionCard(ctx context.Context, card *SongSuggestionCard) (*SongSuggestionCard, error) {
	tx := GetDB().WithContext(ctx).Create(card)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return card, nil
}

func GetSongSuggestionCardByID(ctx context.Context, id int32) (*SongSuggestionCard, error) {
	card := &SongSuggestionCard{}
	tx := GetDB().WithContext(ctx).Where("song_suggestion_card_id = ?", id).First(card)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return card, nil
}

func ListSongSuggestionCards(ctx context.Context, limit int, offset int) ([]*SongSuggestionCard, error) {
	var cards []*SongSuggestionCard
	tx := GetDB().WithContext(ctx).Limit(limit).Offset(offset).Find(&cards)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cards, nil
}

func CreateSongPollCard(ctx context.Context, card *SongPollCard) (*SongPollCard, error) {
	tx := GetDB().WithContext(ctx).Create(card)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return card, nil
}

func GetSongPollCardByID(ctx context.Context, id int32) (*SongPollCard, error) {
	card := &SongPollCard{}
	tx := GetDB().WithContext(ctx).Where("song_poll_card_id = ?", id).First(card)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return card, nil
}

func ListSongPollCards(ctx context.Context, limit int, offset int) ([]*SongPollCard, error) {
	var cards []*SongPollCard
	tx := GetDB().WithContext(ctx).Limit(limit).Offset(offset).Find(&cards)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cards, nil
}

func CreateSongPollCardOption(ctx context.Context, option *SongPollCardOption) (*SongPollCardOption, error) {
	tx := GetDB().WithContext(ctx).Create(option)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return option, nil
}

func GetSongPollCardOptionByID(ctx context.Context, id int32) (*SongPollCardOption, error) {
	opt := &SongPollCardOption{}
	tx := GetDB().WithContext(ctx).Where("song_poll_card_option_id = ?", id).First(opt)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return opt, nil
}

func ListSongPollCardOptionsByPollCardID(ctx context.Context, pollCardID int32) ([]*SongPollCardOption, error) {
	var opts []*SongPollCardOption
	tx := GetDB().WithContext(ctx).Where("song_poll_card_id = ?", pollCardID).Find(&opts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return opts, nil
}
