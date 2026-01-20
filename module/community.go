package module

import (
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math"
	"music-twin-backend/proto/pb"
	"time"
)

func GetRandomCommunityMockHighlightsFromTopPicks(
	ctx context.Context,
	userId int32,
) []*pb.CommunityHighlight {

	topSongs, topArtists := GetUserTopPicks(ctx, userId)
	highlights := []*pb.CommunityHighlight{}

	// 1. Song-like highlight (15–50)
	if len(topSongs) > 0 {
		song := topSongs[int(hashToUint64("song_pick")%uint64(len(topSongs)))]
		songName := safeStr(song.AppleMusicSongName)
		artistName := safeStr(song.AppleMusicArtistName)

		count := stableGrowingNumber(
			"song:"+songName,
			userId,
			15,
			50,
		)

		highlights = append(highlights, &pb.CommunityHighlight{
			Type:           proto.String("song_like"),
			Title:          proto.String(fmt.Sprintf("%d people liked ‘%s’ too", count, songName)),
			Subtitle:       proto.String("More people in your area are vibing to " + artistName),
			AvatarImageUrl: song.AppleMusicSongImageUrl,
		})
	}

	// 2. Artist-like highlight (15–50)
	if len(topArtists) > 0 {
		artist := topArtists[int(hashToUint64("artist_pick")%uint64(len(topArtists)))]
		artistName := safeStr(artist.AppleMusicArtistName)

		count := stableGrowingNumber(
			"artist:"+artistName,
			userId,
			15,
			50,
		)

		highlights = append(highlights, &pb.CommunityHighlight{
			Type:           proto.String("artist_like"),
			Title:          proto.String(fmt.Sprintf("%d users love %s too", count, artistName)),
			Subtitle:       proto.String(artistName + " is trending among MusicTwin users"),
			AvatarImageUrl: artist.AppleMusicArtistImageUrl,
		})
	}

	// 3. Taste overlap highlight (5–20)
	overlapCount := stableGrowingNumber(
		"taste_overlap",
		userId,
		5,
		20,
	)

	highlights = append(highlights, &pb.CommunityHighlight{
		Type:     proto.String("taste_overlap"),
		Title:    proto.String(fmt.Sprintf("%d users have over 80%% music taste overlap with you", overlapCount)),
		Subtitle: proto.String("Matching unlocks as the community grows"),
	})

	return highlights
}

func stableGrowingNumber(
	key string,
	userId int32,
	min int,
	max int,
) int {
	day := currentDayIndex()

	// Base hash
	baseHash := hashToUint64(fmt.Sprintf("%d:%s", userId, key))

	// Base value in range
	rangeSize := max - min + 1
	base := int(baseHash%uint64(rangeSize)) + min

	// Small deterministic daily growth (0–2)
	dailyGrowth := int((baseHash >> 8) % 3)

	value := base + int(day)*dailyGrowth

	// Clamp
	return int(math.Min(float64(value), float64(max)))
}

func hashToUint64(s string) uint64 {
	h := sha1.Sum([]byte(s))
	return binary.BigEndian.Uint64(h[:8])
}

func currentDayIndex() int64 {
	return time.Now().UTC().Unix() / 86400
}

func safeStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
