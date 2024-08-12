package router

import (
	"root/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Post
	post := api.Group("/post")
	post.Get("/", handler.GetPosts)
	post.Get("/:id", handler.GetPost)
	post.Post("/", handler.CreatePost)
	post.Put("/:id", handler.UpdatePost)
	post.Delete("/:id", handler.DeletePost)
}
