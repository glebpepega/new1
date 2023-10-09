package apiserver

import (
	"fmt"
	"os"
	"strconv"

	"github.com/glebpepega/new1/internal/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

type Pagination struct {
	PageVolume int `validate:"required"`
	PageNum    int `validate:"required"`
}

func (api *ApiServer) validateEditMiddleware(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return api.error(c, 400, err)
	}

	n := models.News{}

	if err = c.BodyParser(&n); err != nil {
		return api.error(c, 400, err)
	}

	if err = validate.Struct(&n); err != nil {
		return api.error(c, 400, err)
	}

	n.Id = id

	c.Locals("validatedEdit", n)

	return c.Next()
}

func (api *ApiServer) validateListMiddleware(c *fiber.Ctx) error {
	p := Pagination{}

	if err := c.BodyParser(&p); err != nil {
		return api.error(c, 400, err)
	}

	if err := validate.Struct(&p); err != nil {
		return api.error(c, 400, err)
	}

	c.Locals("validatedPagination", p)

	return c.Next()
}

func (api *ApiServer) authMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	authHeader := headers["Authorization"]

	switch authHeader {
	case os.Getenv("AUTH_KEY"):
		return c.Next()
	case "":
		return api.error(c, 400, fmt.Errorf("missing Authorization header"))
	default:
		return api.error(c, 400, fmt.Errorf("authorization failed"))
	}
}

func (api *ApiServer) logMiddleware(c *fiber.Ctx) error {
	logger := api.Logger.WithFields(logrus.Fields{
		"remote_ip":   c.IP(),
		"remote_port": c.Port(),
	})

	logger.Infof("%s request at %s", c.Method(), c.Request().URI())

	return c.Next()
}
