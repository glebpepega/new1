package apiserver

import (
	"fmt"

	"github.com/glebpepega/new1/internal/models"
	"github.com/gofiber/fiber/v2"
)

type listResponseBody struct {
	Success bool
	News    []models.News
}

func (api *ApiServer) handleEditId(c *fiber.Ctx) error {
	n := c.Locals("validatedEdit").(models.News)
	nc := models.NewsCategories{}
	nc.NewsId = n.Id
	if err := api.db.UpdateNews(&n, &nc); err != nil {
		return api.error(c, 400, err)
	}
	return c.Status(200).Send([]byte(`{"status": "updated"}`))
}

func (api *ApiServer) handleList(c *fiber.Ctx) error {
	p := c.Locals("validatedPagination").(Pagination)
	news, err := api.db.GetAllNews(p.PageNum, p.PageVolume)
	if err != nil {
		return api.error(c, 400, err)
	}
	var body listResponseBody
	if len(news) > 0 {
		body = listResponseBody{
			Success: true,
			News:    news,
		}
	} else {
		body = listResponseBody{
			Success: false,
		}
	}
	return c.Status(200).JSON(body)
}

func (api *ApiServer) error(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).Send([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
}
