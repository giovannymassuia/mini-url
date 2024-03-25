package domain

type Link struct {
	Id       string
	LongUrl  string
	ShortUrl string
}

func NewLink(id, longUrl, shortUrl string) *Link {
	return &Link{
		Id:       id,
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}
}
