package v1

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApiV1(app fiber.Router, db *gorm.DB) {
	api := app.Group("/v1")

	ApiV1Auth(api, db)
	ApiV1User(api, db)
}
