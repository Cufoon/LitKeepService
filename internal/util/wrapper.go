package util

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/constant"
)

type Res struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data any    `json:"data"`
}

func ResOK(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(&Res{
		Code: 0,
		Info: "success",
		Data: data,
	})
}

func ResFail(c *fiber.Ctx, code int, message string) error {
	return c.Status(fiber.StatusOK).JSON(&Res{
		Code: code,
		Info: message,
		Data: nil,
	})
}

func ResFailWithStatusCode(c *fiber.Ctx, status int, code int, message string) error {
	return c.Status(status).JSON(&Res{
		Code: code,
		Info: message,
		Data: nil,
	})
}

func ResFailWithError(c *fiber.Ctx, status int, code int, err error) error {
	return c.Status(status).JSON(&Res{
		Code: code,
		Info: err.Error(),
		Data: nil,
	})
}

func ResBadBody(c *fiber.Ctx) error {
	return ResFailWithError(c, 400, 1, constant.ErrRequestBodyNotValid)
}

type FiberCtx fiber.Ctx

func (c *FiberCtx) ResBadBody() error {
	return ResFailWithError((*fiber.Ctx)(c), 400, 1, constant.ErrRequestBodyNotValid)
}

func Timestamp2Time(timestamp *int64) *time.Time {
	if timestamp == nil {
		return nil
	}
	r := time.UnixMilli(*timestamp)
	return &r
}
