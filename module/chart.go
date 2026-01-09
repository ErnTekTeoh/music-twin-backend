package module

import (
	"context"
	"fmt"
	"music-twin-backend/common"
	"net/url"
)

func GetTop100ChartSongsAndAlbums(ctx context.Context) (*common.AppleMusicChartAPIResponse, error) {
	baseUrl := "https://api.music.apple.com/v1/catalog/sg/charts"
	u, _ := url.Parse(baseUrl)
	q := u.Query()

	q.Set("types", "songs,albums")
	q.Set("chart", "most-played")
	q.Set("limit", "100")
	u.RawQuery = q.Encode()
	fullUrl := u.String()
	fmt.Println(fullUrl)

	resp, err := common.CallHTTPEndpointWithHeaders(fullUrl, "GET", nil)
	if err != nil {
		return nil, err
	}
	return ParseAppleMusicChartResponse(resp), nil
}
