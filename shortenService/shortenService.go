package shortenService

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/subramanian0331/urlShortener/models"
	"github.com/subramanian0331/urlShortener/store"
	"time"
)

type IShortenService interface {
	Expand(shortCode string) (string, error)
	Shorten(longUrl string) (string, error)
}

type shortenService struct {
	db   store.IStorage
	host string
	port string
}

func NewShortenService(db store.IStorage, host string, port string) IShortenService {
	return &shortenService{
		db:   db,
		host: host,
		port: port,
	}
}

func (s *shortenService) Expand(shortCode string) (string, error) {
	url, err := s.db.Get(shortCode)
	if err != nil {
		return "", err
	}
	fmt.Println(url.LongUrl)
	return url.LongUrl, nil
}

func (s *shortenService) Shorten(longUrl string) (string, error) {
	u := new(models.URL)
	u.LongUrl = longUrl
	u.ShortUrl = xid.New().String()
	u.CreatedAt = time.Now()

	// save the data to key value db dynamodb
	err := s.db.Save(u)
	if err != nil {
		return "", err
	}
	shortUrl := fmt.Sprintf("http://%s:%s/%s", s.host, s.port, u.ShortUrl)
	return shortUrl, nil
}
