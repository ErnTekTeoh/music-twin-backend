package common

type AppleMusicSearchResponse struct {
	Results struct {
		Songs struct {
			Data []struct {
				ID         string `json:"id"`
				Attributes struct {
					ArtistName string `json:"artistName"`
					Name       string `json:"name"`
					Artwork    struct {
						Url string `json:"url"`
					} `json:"artwork"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"songs"`
		Artists struct {
			Data []struct {
				ID         string `json:"id"`
				Attributes struct {
					Name    string `json:"name"`
					Artwork struct {
						Url string `json:"url"`
					} `json:"artwork"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"artists"`
	} `json:"results"`
}
