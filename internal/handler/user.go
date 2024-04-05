package handler

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/constant"
	"cufoon.litkeep.service/internal/service"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/internal/vo"
	"cufoon.litkeep.service/pkg/flow"
)

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type UserHandler struct {
	userService *service.UserService
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	data := new(vo.UserRegisterReq)
	err := c.BodyParser(data)
	if err != nil || data.Email == "" || data.Password == "" {
		return util.ResBadBody(c)
	}
	err = uh.userService.Register(&box.UserRegisterData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrAccountExist) {
			return util.ResFail(c, 1, "该邮箱已经注册过")
		}
		if errors.Is(err, constant.ErrAccountIDExist) {
			return util.ResFail(c, 2, "内部id重复，请稍后重试！")
		}
		return err
	}
	token, err := uh.userService.Login(&box.UserLoginData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrLoginEmailNotExist) || errors.Is(err, constant.ErrLoginPassWrong) {
			return util.ResFail(c, 2, "账户或者密码错误")
		}
		return err
	}
	return util.ResOK(c, &vo.UserRegisterRes{
		Signed: true,
		Token:  token,
	})
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	data := new(vo.UserLoginReq)
	err := c.BodyParser(data)
	if err != nil || data.Email == "" || data.Password == "" {
		return util.ResBadBody(c)
	}
	token, err := uh.userService.Login(&box.UserLoginData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrLoginEmailNotExist) || errors.Is(err, constant.ErrLoginPassWrong) {
			return util.ResFail(c, 2, "账户或者密码错误")
		}
		return err
	}
	return util.ResOK(c, &vo.UserLoginRes{
		Logined: true,
		Token:   token,
	})
}

func (uh *UserHandler) GetInfo(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResBadBody(c)
	}
	info, err := uh.userService.GetInfo(userID)
	if err != nil {
		return err
	}
	iconPath := ""
	if info.HasIcon {
		iconPath = "UserIcon/" + info.UserID + "?time=" + info.UpdateTime
	}
	return util.ResOK(c, &vo.UserGetInfoRes{
		NickName: info.NickName,
		UserID:   info.UserID,
		Email:    info.Email,
		IconPath: iconPath,
	})
}

func (uh *UserHandler) ChangeIcon(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)
	fileB, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	userID := flow.GetUserID(c)
	err = uh.userService.SetUserIcon(userID, fileB)
	if err != nil {
		return err
	}
	return util.ResOK(c, &vo.UserChangeNickNameRes{
		Changed: true,
	})
}

func (uh *UserHandler) GetIcon(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	icon, err := uh.userService.GetIcon(userID)
	if err != nil {
		if errors.Is(err, constant.ErrAccountNoIcon) {
			return util.ResFailWithStatusCode(c, 404, 1, err.Error())
		}
		return err
	}
	c.Set("Content-Type", "image/jpeg")
	return c.Send(icon)
}

func (uh *UserHandler) ChangeNickName(c *fiber.Ctx) error {
	data := new(vo.UserChangeNickNameReq)
	err := c.BodyParser(data)
	if err != nil || data.NickName == "" {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	err = uh.userService.SetUserNickName(userID, data.NickName)
	if err != nil {
		return util.ResFail(c, 1, "修改失败")
	}
	return util.ResOK(c, &vo.UserChangeNickNameRes{
		Changed: true,
	})
}
