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

type AppleMusicChartAPIResponse struct {
	Results struct {
		Songs []struct {
			Chart string     `json:"chart"`
			Data  []SongData `json:"data"`
		} `json:"songs"`
		Albums []struct {
			Chart string     `json:"chart"`
			Data  []SongData `json:"data"`
		} `json:"albums"`
	} `json:"results"`
	Meta struct {
		Results struct {
			Order    []string `json:"order"`
			RawOrder []string `json:"rawOrder"`
		} `json:"results"`
	} `json:"meta"`
}

type SongData struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Href       string         `json:"href"`
	Attributes SongAttributes `json:"attributes"`
}

type AlbumData struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Href       string          `json:"href"`
	Attributes AlbumAttributes `json:"attributes"`
}

type AlbumAttributes struct {
	ArtistName          string         `json:"artistName"`
	Artwork             Artwork        `json:"artwork"`
	Copyright           string         `json:"copyright"`
	EditorialNotes      EditorialNotes `json:"editorialNotes"`
	GenreNames          []string       `json:"genreNames"`
	IsCompilation       bool           `json:"isCompilation"`
	IsComplete          bool           `json:"isComplete"`
	IsMasteredForItunes bool           `json:"isMasteredForItunes"`
	IsSingle            bool           `json:"isSingle"`
	Name                string         `json:"name"`
	PlayParams          PlayParams     `json:"playParams"`
	RecordLabel         string         `json:"recordLabel"`
	ReleaseDate         string         `json:"releaseDate"` // YYYY-MM-DD
	TrackCount          int            `json:"trackCount"`
	UPC                 string         `json:"upc"`
	URL                 string         `json:"url"`
}
type SongAttributes struct {
	AlbumName            string     `json:"albumName"`
	ArtistName           string     `json:"artistName"`
	Artwork              Artwork    `json:"artwork"`
	ComposerName         string     `json:"composerName"`
	DiscNumber           int        `json:"discNumber"`
	DurationInMillis     int        `json:"durationInMillis"`
	GenreNames           []string   `json:"genreNames"`
	HasLyrics            bool       `json:"hasLyrics"`
	IsAppleDigitalMaster bool       `json:"isAppleDigitalMaster"`
	Isrc                 string     `json:"isrc"`
	Name                 string     `json:"name"`
	PlayParams           PlayParams `json:"playParams"`
	Previews             []Preview  `json:"previews"`
	ReleaseDate          string     `json:"releaseDate"`
	TrackNumber          int        `json:"trackNumber"`
	Url                  string     `json:"url"`
}

type ArtistData struct {
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Href       string           `json:"href"`
	Attributes ArtistAttributes `json:"attributes"`
}

type ArtistAttributes struct {
	Name    string  `json:"name"`
	Artwork Artwork `json:"artwork"`
	Url     string  `json:"url"`
}

type PlayParams struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

type Preview struct {
	Url string `json:"url"`
}

type Artwork struct {
	BgColor    string `json:"bgColor"`
	Height     int    `json:"height"`
	TextColor1 string `json:"textColor1"`
	TextColor2 string `json:"textColor2"`
	TextColor3 string `json:"textColor3"`
	TextColor4 string `json:"textColor4"`
	Url        string `json:"url"`
	Width      int    `json:"width"`
}

type EditorialNotes struct {
	Short    string `json:"short"`
	Standard string `json:"standard"`
	Tagline  string `json:"tagline"`
}
