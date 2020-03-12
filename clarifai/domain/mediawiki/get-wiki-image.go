package wiki

type MediaWiki struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages struct {
			ID struct {
				Pageid   int
				Ns       int    `json:"ns"`
				Title    string `json:"title"`
				Original struct {
					Source string `json:"source"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"original"`
			}
		} `json:"pages"`
	} `json:"query"`
}