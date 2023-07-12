package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/subramanian0331/urlShortener/shortenService"
	"net/http"
)

type BaseHandler struct {
	svc shortenService.IShortenService
}

func NewBaseHandler(service shortenService.IShortenService) *BaseHandler {

	return &BaseHandler{
		svc: service,
	}
}

func (h *BaseHandler) RedirectUrlHandler(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			fmt.Print(err)
		}
	}()
	shortCode := c.Params("shortCode")
	fmt.Println("shortcode:", shortCode)
	longUrl, err := h.svc.Expand(shortCode)
	if err != nil {
		return err
	}
	return c.Redirect(longUrl, http.StatusMovedPermanently)
}
func (h *BaseHandler) ShortenUrlHandler(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			fmt.Print(err)
		}
	}()
	var url string
	url = string(c.Body())
	shortUrl, err := h.svc.Shorten(url)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(shortUrl)
}
