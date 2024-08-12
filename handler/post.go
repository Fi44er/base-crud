package handler

import (
	"root/db"
	"root/model"

	"github.com/gofiber/fiber/v2"
)

func GetPost(ctx *fiber.Ctx) error {
	posts := db.DB.Find(&[]model.Post{})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": posts})
}

func GetPosts(ctx *fiber.Ctx) error {
	return nil
}

func CreatePost(ctx *fiber.Ctx) error {
	post := new(model.Post)
	if err := ctx.BodyParser(post); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	db.DB.Create(&post)
	return ctx.Status(fiber.StatusOK).JSON(post)
}

func UpdatePost(ctx *fiber.Ctx) error {
	return nil
}

func DeletePost(ctx *fiber.Ctx) error {
	return nil
}
