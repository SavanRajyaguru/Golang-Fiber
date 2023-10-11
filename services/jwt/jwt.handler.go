package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savan/database"
	"github.com/savan/helper"
	model "github.com/savan/services/association-test"
)

func GenerateJwt(c *fiber.Ctx) error {
	var id uint
	query := database.DB
	err := query.Model(&model.Emp{}).Debug().Select("card_id").First(&id, "id = ?", 2).Error
	if err != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	tokenString, err := helper.GenerateJwt(id)
	if err != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "JWT generated unsuccessfully",
			// Data:    tokenString,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "JWT generated successfully",
		Data:    tokenString,
	})
}
