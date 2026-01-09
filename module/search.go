package module

import (
	"context"
	"encoding/json"
	"fmt"
	"music-twin-backend/common"
	"net/url"
)

func SearchAppleMusicTrack(ctx context.Context, title string, artistName string) (*common.AppleMusicSearchResponse, error) {
	baseUrl := "https://api.music.apple.com/v1/catalog/sg/search"
	u, _ := url.Parse(baseUrl)
	q := u.Query()

	// Combine title and artistName for 'term'
	term := ""
	if title != "" && artistName != "" {
		term = fmt.Sprintf("%s %s", title, artistName)
	} else if title != "" {
		term = title
	} else if artistName != "" {
		term = artistName
	}
	q.Set("term", term)
	q.Set("types", "songs")
	u.RawQuery = q.Encode()
	fullUrl := u.String()
	fmt.Println(fullUrl)

	resp, err := common.CallHTTPEndpointWithHeaders(fullUrl, "GET", nil)
	if err != nil {
		return nil, err
	}
	return ParseAppleMusicResponse(resp), nil
}

func SearchAppleMusicArtist(ctx context.Context, artistName string) (*common.AppleMusicSearchResponse, error) {
	baseUrl := "https://api.music.apple.com/v1/catalog/sg/search"
	u, _ := url.Parse(baseUrl)
	q := u.Query()

	// Combine title and artistName for 'term'
	term := ""
	if artistName != "" {
		term = artistName
	}
	q.Set("term", term)
	q.Set("types", "artists")
	u.RawQuery = q.Encode()
	fullUrl := u.String()

	resp, err := common.CallHTTPEndpointWithHeaders(fullUrl, "GET", nil)
	if err != nil {
		return nil, err
	}
	return ParseAppleMusicResponse(resp), nil
}

//
//func SearchArtist(ctx context.Context, artistName string) (*common.DiscogsSearchResponse, error) {
//	baseUrl := "https://api.discogs.com/database/search"
//	u, _ := url.Parse(baseUrl)
//	q := u.Query()
//
//	if artistName != "" {
//		q.Set("query", artistName)
//	}
//	q.Set("type", "artist")
//	u.RawQuery = q.Encode()
//	fullUrl := u.String()
//
//	headers := map[string]string{
//		"Authorization": "Discogs key=siwseMRvWtdyUGJKuDEv, secret=ODUDvwoAzbbCiyTswyGheMiKqMObxAQU",
//		"Cookie":        "__cf_bm=s69EMZXPidl15QC99p0cFDBiH_.YYkMqbPUjF.Q5UHs-1765091907-1.0.1.1-qe6mv36rfxLU2RVyt_MDY6xpj1Wo6Ob.hL3FKJ84C2DVB3b08aUfZcFQSFBntpI1_S8CAEMoK9bgHBK.TkFRSSS4luepMaLW2mtWGT9KuQA",
//	}
//	resp, err := common.CallHTTPEndpointWithHeaders(fullUrl, "GET", headers, nil)
//	if err != nil {
//		return nil, err
//	}
//	return ParseDiscogsResponse(resp), nil
//}
//
//func SearchTrack(ctx context.Context, title string, artistName string) (*common.DiscogsSearchResponse, error) {
//	baseUrl := "https://api.discogs.com/database/search"
//	u, _ := url.Parse(baseUrl)
//	q := u.Query()
//
//	if artistName != "" {
//		q.Set("artist", artistName)
//	}
//	if title != "" {
//		q.Set("title", title)
//	}
//
//	q.Set("type", "release")
//
//	u.RawQuery = q.Encode()
//	fullUrl := u.String()
//
//	headers := map[string]string{
//		"Authorization": "Discogs key=siwseMRvWtdyUGJKuDEv, secret=ODUDvwoAzbbCiyTswyGheMiKqMObxAQU",
//		"Cookie":        "__cf_bm=s69EMZXPidl15QC99p0cFDBiH_.YYkMqbPUjF.Q5UHs-1765091907-1.0.1.1-qe6mv36rfxLU2RVyt_MDY6xpj1Wo6Ob.hL3FKJ84C2DVB3b08aUfZcFQSFBntpI1_S8CAEMoK9bgHBK.TkFRSSS4luepMaLW2mtWGT9KuQA",
//	}
//	resp, err := common.CallHTTPEndpointWithHeaders(fullUrl, "GET", headers, nil)
//	if err != nil {
//		return nil, err
//	}
//	return ParseDiscogsResponse(resp), nil
//}

func ParseAppleMusicResponse(body []byte) *common.AppleMusicSearchResponse {
	var resp *common.AppleMusicSearchResponse

	err := json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("error:", err)
		return resp
	}

	if resp.Results.Songs.Data == nil {
		fmt.Println("No songs found.")
		return resp
	}

	fmt.Printf("Found %d songs\n", len(resp.Results.Songs.Data))

	return resp
}

func ParseAppleMusicChartResponse(body []byte) *common.AppleMusicChartAPIResponse {
	var resp *common.AppleMusicChartAPIResponse

	err := json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("error:", err)
		return resp
	}

	if len(resp.Results.Songs) == 0 || resp.Results.Songs[0].Data == nil {
		fmt.Println("No songs found.")
		return resp
	}

	fmt.Printf("Found %d songs\n", len(resp.Results.Songs[0].Data))

	return resp
}

//
//func ParseDiscogsResponse(body []byte) *common.DiscogsSearchResponse {
//	var resp *common.DiscogsSearchResponse
//
//	err := json.Unmarshal(body, &resp)
//	if err != nil {
//		fmt.Println("error:", err)
//		return resp
//	}
//
//	fmt.Printf("Found %d results\n", len(resp.Results))
//
//	for _, r := range resp.Results {
//		fmt.Println("Title:", r.Title)
//		fmt.Println("Type:", r.Type)
//		fmt.Println("Resource URL:", r.ResourceURL)
//		fmt.Println()
//	}
//	return resp
//}
