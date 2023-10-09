package apiserver

import (
	"os"

	"github.com/glebpepega/new1/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	fiber  *fiber.App
	db     *db.DB
	logger *logrus.Logger
}

func New() *ApiServer {
	return &ApiServer{}
}

func (api *ApiServer) configure() {
	api.fiber = fiber.New()
	api.logger = logrus.New()
	api.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	api.db = db.New()
	if err := api.db.Init(); err != nil {
		api.logger.Fatal(err)
	}
}

func (api *ApiServer) Start() {
	api.configure()
	api.logger.Info("starting api server")
	api.fiber.Use(api.authMiddleware, api.logMiddleware)
	api.fiber.Use("/edit/:id", api.validateEditMiddleware)
	api.fiber.Use("/list", api.validateListMiddleware)
	api.fiber.Post("/edit/:id", api.handleEditId)
	api.fiber.Get("/list", api.handleList)
	api.logger.Fatal(api.fiber.Listen(os.Getenv("LISTENING_PORT")))
}
