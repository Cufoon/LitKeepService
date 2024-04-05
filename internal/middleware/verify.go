package middleware

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/util"
)

func Validate(s any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.BodyParser(s)
		if err != nil {
			return util.ResFailWithStatusCode(c, 400, 1, "请求参数错误")
		}
		return c.Next()
	}
}
