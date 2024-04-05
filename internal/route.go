package internal

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/handler"
	"cufoon.litkeep.service/internal/middleware"
)

func InitRoute(app *fiber.App, c *handler.Handler, m *middleware.MiddleWare) {
	v1 := app.Group("/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1.Get("/token_verify", m.Auth, c.TokenHandler.Verify)
	v1.Post("/version_android", c.OtherHandler.CheckAndroidAppUpdate)

	v1.Get("/user_info", m.Auth, c.UserHandler.GetInfo)
	v1.Get("/user_icon", m.Auth, c.UserHandler.GetIcon)

	v1.Post("/user_login", c.UserHandler.Login)
	v1.Post("/user_register", c.UserHandler.Register)
	v1.Post("/user_info_change_nickname", m.Auth, c.UserHandler.ChangeNickName)
	v1.Post("/user_info_change_icon", m.Auth, c.UserHandler.ChangeIcon)

	v1.Post("/kind_create", m.Auth, c.BillKindHandler.Create)
	v1.Post("/kind_query", m.Auth, c.BillKindHandler.Query)
	v1.Post("/kind_update", m.Auth, c.BillKindHandler.Modify)
	v1.Post("/kind_delete", m.Auth, c.BillKindHandler.Delete)

	v1.Post("/bill_create", m.Auth, c.BillRecordHandler.Create)
	v1.Post("/bill_query", m.Auth, c.BillRecordHandler.Query)
	v1.Post("/bill_query_with_kind", m.Auth, c.BillRecordHandler.QueryWithKind)
	v1.Post("/bill_query_page", m.Auth, c.BillRecordHandler.QueryPage)
	v1.Post("/bill_query_page_info", m.Auth, c.BillRecordHandler.QueryPageData)
	v1.Post("/bill_update", m.Auth, c.BillRecordHandler.Modify)
	v1.Post("/bill_delete", m.Auth, c.BillRecordHandler.Delete)
	v1.Post("/bill_query_statistic_day", m.Auth, c.BillRecordHandler.QueryStatisticsDay)
}
