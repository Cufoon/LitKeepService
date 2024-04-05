package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"

	"cufoon.litkeep.service/conf"
	"cufoon.litkeep.service/internal/db/cache"
	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/entity"
	"cufoon.litkeep.service/internal/handler"
	"cufoon.litkeep.service/internal/middleware"
	"cufoon.litkeep.service/internal/service"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/pkg/db"
	"cufoon.litkeep.service/pkg/jwt"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func StartAPP(path string) {
	c, err := conf.NewConf(path)
	if err != nil {
		println(err.Error())
		return
	}
	err = jwt.Init("LT", c.AESKey, c.ED25519KeyPri, c.ED25519KeyPub)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	app := fiber.New(fiber.Config{
		Prefork:       false,
		ServerHeader:  "Cufoon.LitKeep.Server",
		CaseSensitive: true,
		AppName:       "LitKeep",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		//ReduceMemoryUsage: false,
		//Immutable:         false,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			errID := ""
			errInfo := ""
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				errInfo = err.Error()
			} else {
				errID = util.GenerateBytes(16)
				errInfo = strings.Join([]string{"some error happened, please see error", errID}, " ")
			}
			err2 := util.ResFailWithStatusCode(ctx, code, 1, errInfo)
			if err2 != nil {
				return util.ResFailWithStatusCode(ctx, fiber.StatusInternalServerError, 1, errInfo)
			}
			if errID != "" {
				fmt.Println(errID, err.Error())
			}
			return nil
		},
	})

	mariaDB, err := db.NewDB(c)
	if err != nil {
		println(err.Error())
		return
	}

	userDAO := dao.NewUserDAO(mariaDB, &entity.User{})
	billKindDAO := dao.NewBillKindDAO(mariaDB, &entity.BillKind{})
	billRecordDAO := dao.NewBillRecordDAO(mariaDB, &entity.BillRecord{})

	billKindCache := cache.NewSystemBillKindCache(billKindDAO)
	cacheInited := billKindCache.InitCache()
	if cacheInited {
		println("<-----cache init success----->")
	} else {
		println("<-----cache init failed----->")
	}

	userService := service.NewUserService(userDAO, c)
	billKindService := service.NewBillKindService(billKindDAO, billKindCache)
	billRecordService := service.NewBillRecordService(billRecordDAO, billKindDAO)

	userHandler := handler.NewUserHandler(userService)
	billKindHandler := handler.NewBillKindHandler(billKindService)
	billRecordHandler := handler.NewBillRecordHandler(billRecordService, billKindService)

	tokenHandler := handler.NewTokenHandler()
	otherHandler := handler.NewOtherHandler()

	middleWare := middleware.NewMiddleWare()

	InitRoute(app, &handler.Handler{
		UserHandler:       userHandler,
		BillKindHandler:   billKindHandler,
		BillRecordHandler: billRecordHandler,
		TokenHandler:      tokenHandler,
		OtherHandler:      otherHandler,
	}, middleWare)

	err = app.Listen(c.Server.Port)
	if err != nil {
		println(err.Error())
	}
}
