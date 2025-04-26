package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"yuyuid.id/config"
	"yuyuid.id/internal/model"
	"yuyuid.id/internal/repository"
	"yuyuid.id/internal/service"
	"yuyuid.id/internal/util"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func AuthLogin(c *fiber.Ctx, db *gorm.DB) error {
	var req model.LoginUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseError("Invalid Request"))
	}

	err := validate.Struct(req)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, e.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseValidator(validationErrors, "Error Validation"))
	}

	var user model.User
	if err := db.Where("email = ? AND deleted_at IS NULL", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(config.ResponseError("Invalid Credentials"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(config.ResponseError("Invalid credentials"))
	}
	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(config.ResponseError("Could not login"))
	}

	return c.JSON(config.ResponseSuccess(fiber.Map{"token": token}))
}

func AuthRegister(c *fiber.Ctx, db *gorm.DB) error {
	var req model.RegisterUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseError("Invalid JSON format"))
	}

	err := validate.Struct(req)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, e.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseValidator(validationErrors, "Error Validation"))
	}

	userRepo := repository.UserRepositoryImpl{DB: db}
	userService := service.NewUserService(userRepo)

	err = userService.IsUserExist(c.Context(), req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(config.ResponseError("Email has already registered"))
	}

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 4)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(config.ResponseError("Failed to has password"))
	}
	user := model.User{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Status:    "INACTIVE",
	}
	user.Password = string(hashedPassword)

	if err := userService.SaveUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(config.ResponseError("Error Saving User"))
	}

	return c.Status(fiber.StatusCreated).JSON(config.ResponseCreateSuccess(user, "User Registered successfully"))
}
