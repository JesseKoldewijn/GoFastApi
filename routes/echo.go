package routes

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func spritBySlash(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			result += " "
		} else {
			result += string(s[i])
		}
	}

	decodedResult, _ := url.QueryUnescape(result)
	decodedResult = strings.ReplaceAll(decodedResult, "%20", " ")

	return decodedResult
}

func Echo(App *fiber.App) {
	App.Get("/echo/**", func(c fiber.Ctx) error {
		url := c.Params("*")       // Get wildcard values

		if url == "" {
			c.SendStatus(400)
			return c.JSON(fiber.Map{
				"error":   "No message provided",
				"message": "Please provide a message to echo by adding it to the url",
				"example": "http://localhost:8080/echo/Hello%20Mom",
			})
		}


		return c.JSON(fiber.Map{
			"message": spritBySlash(url),
		})
	})
}
