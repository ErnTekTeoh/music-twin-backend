package common

type DiscogsSearchResponse struct {
	Pagination Pagination     `json:"pagination"`
	Results    []SearchResult `json:"results"`
}

type Pagination struct {
	PerPage int            `json:"per_page"`
	Pages   int            `json:"pages"`
	Page    int            `json:"page"`
	Urls    PaginationUrls `json:"urls"`
	Items   int            `json:"items"`
}

type PaginationUrls struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type SearchResult struct {
	Style       []string      `json:"style"`
	Thumb       string        `json:"thumb"`
	Title       string        `json:"title"`
	Country     string        `json:"country"`
	Format      []string      `json:"format"`
	URI         string        `json:"uri"`
	Community   CommunityInfo `json:"community"`
	Label       []string      `json:"label"`
	CatNo       string        `json:"catno"`
	Year        string        `json:"year"`
	Genre       []string      `json:"genre"`
	ResourceURL string        `json:"resource_url"`
	Type        string        `json:"type"`
	ID          int           `json:"id"`

	// Optional fields that appear only sometimes
	Barcode []string `json:"barcode"`
}

type CommunityInfo struct {
	Want int `json:"want"`
	Have int `json:"have"`
}
