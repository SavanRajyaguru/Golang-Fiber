package associationtest

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/savan/database"
	"github.com/savan/helper"
)

func GetEmp(c *fiber.Ctx) error {

	// var res []SendCardResponse
	var res []Card

	data := database.DB.Model(&Card{}).Preload("Employees").Find(&res).Error
	// data := database.DB.Model(&Emp{}).Preload("CardDetails").Find(&res).Error
	// data := database.DB.Model(&Emp{}).Preload("Employees").Find(&res).Error

	if data != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to load data",
			Data:    data,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Success for the Emp",
		Data:    res,
	})
}

func TransactionDemo(c *fiber.Ctx) error {
	type getData struct {
		Name     string `validate:"required,min=2" json:"name"`
		CardName string `validate:"required" json:"cardname"`
	}
	var body getData
	if err := c.BodyParser(&body); err != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(&body); err != nil {
		return helper.SendResponse(c, helper.Response{
			Message: "Validation failed",
			Code:    fiber.StatusBadRequest,
			Data:    strings.Split(err.Error(), "\n"),
		})
	}

	// if err := c.BodyParser(&body); err != nil {

	// 	return helper.SendResponse(c, helper.Response{
	// 		Code:    fiber.StatusBadRequest,
	// 		Message: err.Error(),
	// 		Data:    nil,
	// 	})
	// }

	tx := database.DB.Begin()
	tx = tx.Set("gorm:query_option", "SET TRANSACTION ISOLATION LEVEL READ COMMITTED;")

	err := tx.Model(&Emp{}).Create(&Emp{Name: "Hello1"}).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("hello1")
		return err
	}
	var name string
	err = tx.Model(&Emp{}).Create(&Emp{Name: "Hello5"}).Select("name").Last(&name).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("hello3")
		return err
	}

	if err = tx.Commit().Error; err != nil {
		fmt.Println("Error committing transaction:", err)
	}

	// tx := database.DB.Transaction(func(tx *gorm.DB) error {
	// 	var id int
	// 	err := tx.Model(&Card{}).Select("id").Where("name = ?", body.CardName).First(&id).Error
	// 	if err != nil {
	// 		fmt.Println("First", err)
	// 		return err
	// 	}
	// 	fmt.Println("ID>>>", id)

	// 	newEmp := Emp{
	// 		Name: body.Name,
	// 	}

	// 	err = tx.Model(&Emp{}).Create(&newEmp).Error
	// 	if err != nil {
	// 		fmt.Println("Second", err)
	// 		return err
	// 	}

	// 	err = tx.Model(&newEmp).Association("CardDetails").Append(&Card{ID: uint(id)})
	// 	if err != nil {
	// 		fmt.Println(">>>>>", err)
	// 		return err
	// 	}
	// 	return nil
	// })

	// if tx != nil {
	// 	return helper.SendResponse(c, helper.Response{
	// 		Code:    fiber.StatusBadRequest,
	// 		Message: "Error in transaction",
	// 		Data:    tx,
	// 	})
	// }

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Success for the TransactionDemo",
		Data:    name,
	})
}
