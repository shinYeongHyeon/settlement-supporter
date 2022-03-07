package coreError

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func NotFoundError(context *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello ✋\nWe are Settlement Supporter API.\nRequest URL %s is NOT FOUND", context.OriginalURL())

	return context.Status(fiber.StatusNotFound).SendString(msg)
}
