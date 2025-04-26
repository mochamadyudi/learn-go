package v1

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"yuyuid.id/internal/handlers"
)

func ApiV1User(router fiber.Router, db *gorm.DB) {
	api := router.Group("/user")

	api.Get("", func(c *fiber.Ctx) error {
		return handlers.UserList(c, db)
	})
}
