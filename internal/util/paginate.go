package util

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func Paginate(c *fiber.Ctx) Pagination {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}
