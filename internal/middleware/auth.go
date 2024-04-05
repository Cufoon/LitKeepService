package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/constant"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/pkg/flow"
	"cufoon.litkeep.service/pkg/jwt"
)

func (mw *MiddleWare) Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization", "")
	if token == "" {
		return c.SendStatus(401)
	}
	data, err := jwt.Parse(token)
	if err != nil {
		return c.SendStatus(401)
	}
	if data.ExpireTime <= time.Now().UnixMicro() {
		return util.ResFailWithError(c, 401, 2, constant.ErrAccountTokenExpired)
	}
	err = flow.SetUserID(c, data.UserId)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.Next()
}
