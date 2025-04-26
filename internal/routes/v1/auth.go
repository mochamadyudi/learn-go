package v1

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"yuyuid.id/internal/handlers"
)

func ApiV1Auth(router fiber.Router, db *gorm.DB) {
	api := router.Group("/auth")

	api.Post("/login", func(c *fiber.Ctx) error {
		return handlers.AuthLogin(c, db)
	})
	api.Post("/register", func(c *fiber.Ctx) error {
		return handlers.AuthRegister(c, db)
	})
}
