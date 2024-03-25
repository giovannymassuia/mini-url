package dto

type CrateShortLinkRequest struct {
	LongUrl string `json:"long_url"`
}

type CreateShortLinkResponse struct {
	ShortUrl string `json:"short_url"`
}
