package handler

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/internal/vo"
)

func NewOtherHandler() *OtherHandler {
	return &OtherHandler{}
}

type OtherHandler struct{}

func (o *OtherHandler) CheckAndroidAppUpdate(c *fiber.Ctx) error {
	data := new(vo.AndroidAppUpdateCheckReq)
	err := c.BodyParser(data)
	if err != nil || data.Now == 0 {
		return util.ResBadBody(c)
	}
	if data.Now < 30000001 {
		return util.ResOK(c, &vo.AndroidAppUpdateCheckRes{
			Update: true,
			URL:    "https://cufoon.com",
		})
	}
	return util.ResOK(c, &vo.AndroidAppUpdateCheckRes{
		Update: false,
	})
}
