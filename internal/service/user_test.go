package service

import (
	"testing"

	"cufoon.litkeep.service/conf"
	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/entity"
	"cufoon.litkeep.service/pkg/db"
)

func TestCreateUser(t *testing.T) {
	gc, err := conf.NewConf("../../dev.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	mariaDB, err := db.NewDB(gc)
	if err != nil {
		t.Error(err)
		return
	}
	userDAO := dao.NewUserDAO(mariaDB, &entity.User{})
	userService := NewUserService(userDAO, gc)
	err = userService.Register(&box.UserRegisterData{
		Email:    "cufoon@gmail.com",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
}
