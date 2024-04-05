package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"cufoon.litkeep.service/conf"
	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/constant"
	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/pkg/jwt"
)

func NewUserService(dao *dao.UserDAO, config *conf.Conf) *UserService {
	return &UserService{dao: dao, config: config}
}

type UserService struct {
	dao    *dao.UserDAO
	config *conf.Conf
}

func (us *UserService) Register(userInfo *box.UserRegisterData) error {
	_, err := us.dao.QueryByEmail(userInfo.Email)
	if err == nil {
		return constant.ErrAccountExist
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	newID := util.Generate8Bytes()
	_, err = us.dao.QueryByUserID(newID)
	if err == nil {
		return constant.ErrAccountIDExist
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err1 := us.dao.Create(&dto.UserCreateData{
			UserID:   newID,
			Email:    userInfo.Email,
			Password: userInfo.Password,
		})
		if err1 != nil {
			return err1
		}
	} else {
		return err
	}
	return nil
}

func (us *UserService) Token(id string, sid uint8, long time.Duration) (string, error) {
	now := time.Now()
	return jwt.Token(&jwt.TokenProperty{
		UserId:     id,
		SessionId:  sid,
		SignedTime: now.UnixMicro(),
		ExpireTime: now.Add(long).UnixMicro(),
	})
}

func (us *UserService) Login(info *box.UserLoginData) (string, error) {
	user, err := us.dao.QueryByEmail(info.Email)
	if err == nil {
		if user.Password == info.Password {
			token, err1 := us.Token(user.UserID, 0, time.Second*time.Duration(us.config.Expire))
			if err1 != nil {
				return "", err1
			}
			return token, nil
		} else {
			return "", constant.ErrLoginPassWrong
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", constant.ErrLoginEmailNotExist
	}
	return "", err
}

func (us *UserService) GetInfo(userID string) (*box.UserInfo, error) {
	user, err := us.dao.QueryByUserID(userID)
	if err == nil {
		hasIcon := true
		if user.Icon == nil {
			hasIcon = false
		}
		return &box.UserInfo{
			NickName:   user.NickName,
			UserID:     user.UserID,
			Email:      user.Email,
			HasIcon:    hasIcon,
			UpdateTime: user.UpdatedAt.Local().String(),
		}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrLoginEmailNotExist
	}
	return nil, err
}

func (us *UserService) GetIcon(userID string) ([]byte, error) {
	user, err := us.dao.QueryByUserID(userID)
	if err == nil {
		if len(user.Icon) == 0 {
			return nil, constant.ErrAccountNoIcon
		}
		return user.Icon, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrLoginEmailNotExist
	}
	return nil, err
}

func (us *UserService) SetUserNickName(userID string, nickName string) error {
	return us.dao.ModifyNickNameByUserID(&dto.UserChangeNickNameData{
		UserID:   userID,
		NickName: nickName,
	})
}

func (us *UserService) SetUserIcon(userID string, icon []byte) error {
	return us.dao.ModifyIconByUserID(&dto.UserChangeIconData{
		UserID: userID,
		Icon:   icon,
	})
}
