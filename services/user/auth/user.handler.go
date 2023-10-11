package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/savan/database"
	"github.com/savan/helper"

	"gorm.io/gorm"
)

type SignupResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type DemoStruct struct {
	// Email string `json:"email"`
	Count int `json:"count"`
}

func SignUpUser(c *fiber.Ctx) error {
	var body UserModel

	if err := c.BodyParser(&body); err != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// convert hex with sha256
	h := sha256.New()
	h.Write([]byte(body.Password))
	body.Password = hex.EncodeToString(h.Sum(nil))

	// check if user is exists or not
	var user SignupResponse
	findUser := database.DB.Model(&user).Limit(1).First(&user, "email = ?", body.Email).Error

	if findUser != nil && findUser.Error() != "record not found" {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Something went wrong",
			// Data:    nil,
		})
	}

	log.Printf("data1: %v, data2: %v", user.Email, body.Email)
	if user.Email == body.Email {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Email already exists!",
			// Data:    nil,
		})
	}

	// insert data to the database
	result := database.DB.Create(&UserModel{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}).Error

	if result != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Signup Failed!",
			Data:    nil,
		})
	}
	findingUserError := database.DB.Model(&UserModel{}).Limit(1).Last(&user).Error

	if findingUserError != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Data not get",
			Data:    nil,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Message: "Sign up Successful",
		Code:    fiber.StatusOK,
		Data:    user,
	})
}

func Query1(db *gorm.DB) *gorm.DB {
	return db.Where("created_at < ?", time.Now())
}
func Query2(db *gorm.DB) *gorm.DB {
	return db.Where("email > ?", "test4@gmail.coms")
}
func GetUser(c *fiber.Ctx) error {
	// var user SignupResponse
	// if err := c.BodyParser(&model.User{}); err != nil {
	// 	return helper.SendResponse(c, helper.Response{
	// 		Code:    fiber.StatusBadRequest,
	// 		Message: "Body is wrong!",
	// 	})
	// }

	var userData []SignupResponse
	var count int64
	// fmt.Println("TIME>>>", currtime)
	// query := database.DB.Model(&model.User{}).Select("email, count(*) count").Group("email").Having("count = ?", 5)
	// data1 := database.DB.Model(&model.User{}).Where("email = ?", query.Find(struct{ email string }{})).Find(&userData).Error
	// data := database.DB.Model(&model.User{}).Select("email, count(*) count").Group("email").Having("count > ?", 5).Find(&userData).Error
	data := database.DB.Model(&UserModel{}).Scopes(Query1, Query2).Find(&userData).Error
	database.DB.Model(&UserModel{}).Count(&count)
	// var userData []DemoStruct
	// data := database.DB.Model(&model.User{}).Select("date(created_at) as date, count(*) as total").Group("date(created_at)").Scan(&userData)

	if data != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Data not get",
			Data:    data,
		})
	}
	d := make(map[string]interface{})
	d["count"] = count
	d["userData"] = userData
	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Data fetch success",
		Data:    d,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// show error
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Id should be integer",
			Data:    err,
		})
	}

	// perform update op
	var userData SignupResponse

	// data := database.DB.Model(&UserModel{}).Where("id = ?", id).Limit(1).First(&userData).Error
	// userData.Email = "test777@gmail.com"
	// database.DB.Table("users").Save(&userData)
	data := database.DB.Model(&UserModel{}).Where("id = ?", id).Update("email", "asdfasdf@adsadsf.com").First(&userData).Error
	if data != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Something went wrong!!",
			// Data:    ,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Data fetch success",
		Data:    userData,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	_, err := c.ParamsInt("id")
	if err != nil {
		// show error
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Id should be integer",
			Data:    err,
		})
	}

	// data := database.DB.Delete(&UserModel{}, id).Error
	var user []SignupResponse
	data := database.DB.Model(&UserModel{}).Unscoped().Where("deleted_at is not null").Find(&user).Error

	if data != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Something went wrong!!",
			// Data:    data,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Delete success",
		Data:    user,
	})
}

func AssociationDemo(c *fiber.Ctx) error {
	type Test struct {
		ID uint `json:"id"`
		// Username  string  `json:"username" gorm:"not null;size:55" validate:"required"`
		// Email     string  `json:"email" gorm:"not null" validate:"required,email"`
		// Password  string  `json:"password" gorm:"not null" validate:"required,gte=6"`
		// CompanyID uint    `json:"-"`
		Company Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	}

	var users Test
	err := database.DB.Model(&Company{}).Preload("Company").Find(&users).Error

	if err != nil {
		return helper.SendResponse(c, helper.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Something wrong",
			Data:    err,
		})
	}

	return helper.SendResponse(c, helper.Response{
		Code:    fiber.StatusOK,
		Message: "Delete success",
		Data:    users,
	})
}
