package handler

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/internal/vo"
)

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

type TokenHandler struct {
}

func (th *TokenHandler) Verify(c *fiber.Ctx) error {
	return util.ResOK(c, &vo.TokenVerifyRes{Verified: true})
}
