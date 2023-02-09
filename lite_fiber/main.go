package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type User struct {
	Name      string
	Phone     string
	Age       int64
	CreatedAt time.Time
}

func main() {
	user := User{
		Name:      "Vasya",
		Phone:     "+380994587834",
		Age:       19,
		CreatedAt: time.Now(),
	}

	router := fiber.New()
	router.Get("/v1/status", GetStatus)
	router.Get("/v1/user", user.GetUser)
	router.Put("v1/user/:name/:phone/:age", user.UpdateUser)

	_ = router.Listen(":8888")
}

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.SendString("status ok")
}

//
//func CreateUser(ctx *fiber.Ctx) error {
//
//	if err != nil {
//		return ctx.SendStatus(fiber.StatusBadRequest)
//	}
//}
//
//func DeleteUser(ctx *fiber.Ctx) error {
//
//	if err != nil {
//		return ctx.SendStatus(fiber.StatusBadRequest)
//	}
//}

func (u *User) UpdateUser(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	phone := ctx.Params("phone")
	age, err := ctx.ParamsInt("age")

	u.Name = name
	u.Phone = phone
	u.Age = int64(age)

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (u *User) GetUser(ctx *fiber.Ctx) error {
	time := u.CreatedAt.String()
	age := strconv.Itoa(int(u.Age))
	return ctx.SendString("Name: " + u.Name + ", phone: " + u.Phone + ", age: " + age + ", has been created: " + time)
}
