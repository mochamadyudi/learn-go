package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"yuyuid.id/config"
	"yuyuid.id/internal/model"
	"yuyuid.id/internal/util"
)

//var validate *validator.Validate

func init() {
	validate = validator.New()
}

func UserList(c *fiber.Ctx, db *gorm.DB) error {
	pagination := util.Paginate(c)
	var users []model.User
	var total int64

	if err := db.Model(&model.User{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseError("Failed to get data"))
	}

	if err := db.Model(&model.User{}).Limit(pagination.Limit).Offset(pagination.Offset).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(config.ResponseError("Failed to fetch users"))
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = total / int64(pagination.Limit)
		if total%int64(pagination.Limit) > 0 {
			totalPages++
		}
	}

	return c.JSON(config.ResponseWithPagination[[]model.User](users, config.Pagination{
		Page:    pagination.Page,
		Limit:   pagination.Limit,
		Total:   int(total),
		Maxpage: int(totalPages),
	}))
}
