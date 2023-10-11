package user

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savan/helper"
)

func SignUpUser(ctx *fiber.Ctx) error {
	var body UserSignUp

	isSuccess := helper.Validate(ctx, &body)

	if !isSuccess {
		return nil
	}

	return helper.SendResponse(ctx, helper.Response{
		Message: "Sign up Successful",
		Code:    200,
		Data:    nil,
	})
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Test(ctx *fiber.Ctx) error {

	ctx.Append("Auth", "jwttoken")
	ctx.Set("Auth1", "savan")
	fmt.Println("data", string(ctx.BodyRaw()))

	data1 := ctx.Locals("User")
	log.Println("asdfasdf>>>", data1)

	data := ctx.GetReqHeaders()

	fmt.Println("Dta>>>", data["Auths"])

	u := new(UserAuth)

	if err := ctx.BodyParser(u); err != nil {
		return err
	}
	fmt.Println("Data: ", u.Username)
	fmt.Println("Data: ", u.Password)

	return helper.SendResponse(ctx, helper.Response{
		Message: "This is test",
		Code:    200,
	})
}
