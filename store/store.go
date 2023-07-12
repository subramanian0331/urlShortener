package store

import "github.com/subramanian0331/urlShortener/models"

type IStorage interface {
	Save(url *models.URL) error
	Get(shortUrl string) (*models.URL, error)
}

func NewStorage(host string) IStorage {
	return &dynamoDB{}
}
