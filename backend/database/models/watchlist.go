package models

type WatchListRow struct {
	ID                int64  `json:"id"`
	Title             string `json:"title"`
	MediaType         int    `json:"mediaType"`
	Genre             string `json:"genre"`
	StreamingPlatform string `json:"streamingPlatform"`
	ShortNote         string `json:"shortNote"`
}
