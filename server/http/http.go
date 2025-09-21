package http

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/BoomTHDev/tattoo_port/config"
	"github.com/BoomTHDev/tattoo_port/databases"
	"github.com/BoomTHDev/tattoo_port/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app  *fiber.App
	db   databases.Database
	conf *config.Config
}

var (
	once         sync.Once
	httpInstance *fiberServer
)

func NewFiberServer(conf *config.Config, db databases.Database) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:    conf.Server.BodyLimit,
		IdleTimeout:  time.Second * time.Duration(conf.Server.TimeOut),
		ErrorHandler: middleware.ErrorHandler(),
	})

	once.Do(func() {
		httpInstance = &fiberServer{
			app:  fiberApp,
			db:   db,
			conf: conf,
		}
	})

	return httpInstance
}

func (h *fiberServer) setupRoutes() {
	h.app.Use(logger.New())
	h.app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(h.conf.Server.AllowOrigins, ","),
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	h.initTattooRouter()

	h.app.Get("/health-check", h.healthCheck)
	h.app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Sorry, endpoint %s %s not found", ctx.Method(), ctx.Path()),
		})
	})
}

func (h *fiberServer) Start() {
	h.setupRoutes()
	h.httpListening()
}

func (h *fiberServer) httpListening() {
	url := fmt.Sprintf(":%d", h.conf.Server.Port)
	if err := h.app.Listen(url); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func (h *fiberServer) healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
