package apiserver

import (
	"os"

	"github.com/glebpepega/new1/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	Fiber  *fiber.App
	Db     *db.DB
	Logger *logrus.Logger
}

func New(db *db.DB, logger *logrus.Logger) *ApiServer {
	return &ApiServer{
		Db:     db,
		Logger: logger,
		Fiber:  fiber.New(),
	}
}

func (api *ApiServer) ConfigureServer() {
	api.Logger.Info("starting api server")
	api.Fiber.Use(api.authMiddleware, api.logMiddleware)
	api.Fiber.Use("/edit/:id", api.validateEditMiddleware)
	api.Fiber.Use("/list", api.validateListMiddleware)
	api.Fiber.Post("/edit/:id", api.handleEditId)
	api.Fiber.Get("/list", api.handleList)
	api.Logger.Error(api.Fiber.Listen(os.Getenv("LISTENING_PORT")))
}
