package main

import (
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/rngallen/fiber-email/config"
	"github.com/rngallen/fiber-email/routers"
)

func main() {

	serverPort := config.Config("SERVER.PORT")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
		CaseSensitive: true,
		ServerHeader:  "Email",
		AppName:       "Email v1.0.0",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	app.Use(helmet.New())

	// Limiter
	app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwred-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.FixedWindow{},
	}))

	// Logs
	file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("log file was not found %s", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "[${time}] ${ip} - ${protocol} - ${status} - ${latency} ${method} ${url}\n",
		TimeFormat: "02-01-2006 15:04:05",
	}))

	app.Server().MaxConnsPerIP = 1

	routers.SetupRouters(app)

	log.Fatal(app.Listen(serverPort))

}
