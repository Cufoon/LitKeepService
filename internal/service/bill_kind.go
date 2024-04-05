package service

import (
	"errors"
	"sort"

	"gorm.io/gorm"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/constant"
	"cufoon.litkeep.service/internal/db/cache"
	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/util"
)

func NewBillKindService(billKindDAO *dao.BillKindDAO, billKindCache *cache.SystemBillKindCache) *BillKindService {
	return &BillKindService{dao: billKindDAO, cache: billKindCache}
}

type BillKindService struct {
	dao   *dao.BillKindDAO
	cache *cache.SystemBillKindCache
}

func (brs *BillKindService) generateKindID() string {
	newKindID := util.GenerateBytes(16)
	for i := 0; i < 10; i++ {
		if !brs.dao.ExistByKindID(newKindID) {
			break
		}
		newKindID = util.GenerateBytes(16)
	}
	return newKindID
}

func (brs *BillKindService) Create(record *box.BillKindCreate) error {
	kindID := brs.generateKindID()

	if record.UpKind != "" && !brs.dao.ExistByUserIDAndKindID(record.UserID, record.UpKind) {
		return constant.ErrBillKindFatherNotExist
	}

	return brs.dao.Create(&dto.BillKindModifyData{
		UserID:      record.UserID,
		KindID:      kindID,
		Name:        record.Name,
		Description: record.Description,
		UpKind:      record.UpKind,
	})
}

// QueryLinear 直接返回一层类型列表，不处理他们的父子关系
func (brs *BillKindService) QueryLinear(userID string) ([]*box.BillKind, error) {
	userKinds, err := brs.dao.QueryByUserIDWithDeletedOver(userID)
	if err != nil {
		return nil, err
	}
	userKinds = brs.cache.JoinSystemBillKind(userKinds)
	r := make([]*box.BillKind, len(userKinds))
	for idx, item := range userKinds {
		r[idx] = &box.BillKind{
			KindID:      item.KindID,
			Name:        item.Name,
			Description: item.Description,
		}
	}
	return r, nil
}

// Query 查询所有类型，返回的是包含父子关系的类型列表
func (brs *BillKindService) Query(userID string) ([]*box.BillKindTree, error) {
	userKinds, err := brs.dao.QueryByUserIDWithDeletedOver(userID)
	if err != nil {
		return nil, err
	}
	userKinds = brs.cache.JoinSystemBillKind(userKinds)

	remap := generateMap(userKinds)
	keys := getKeys(remap)
	sort.Strings(keys)
	r := make([]*box.BillKindTree, len(keys))
	for index, value := range keys {
		r[index] = remap[value]
	}
	return r, nil
}

func (brs *BillKindService) Modify(userID string, billKind *box.BillKind) error {
	modifyData := &dto.BillKindModifyData{
		UserID:      userID,
		KindID:      billKind.KindID,
		Name:        billKind.Name,
		Description: billKind.Description,
		OverKind:    billKind.OverKind,
		UpKind:      billKind.UpKind,
	}
	if brs.cache.ExistByKindID(billKind.KindID) {
		overKind, err := brs.dao.QueryByOverKindIDAndUserIDContainDeleted(userID, billKind.KindID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				modifyData.OverKind = billKind.KindID
				modifyData.KindID = brs.generateKindID()
				return brs.dao.Create(modifyData)
			}
			return err
		}
		if (overKind.DeletedAt != gorm.DeletedAt{}) {
			return constant.ErrBillKindNotExist
		}
		modifyData.KindID = overKind.KindID
	}
	if brs.dao.ExistByUserIDAndKindID(userID, modifyData.KindID) {
		return brs.dao.Modify(modifyData)
	}
	return constant.ErrBillKindNotExist
}

func (brs *BillKindService) Delete(userID, kindID string) error {
	kindID2delete := kindID
	if brs.cache.ExistByKindID(kindID) {
		overKind, err := brs.dao.QueryByOverKindIDAndUserIDContainDeleted(userID, kindID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				systemKind := brs.cache.GetCache()[kindID]
				newKindID := brs.generateKindID()
				isCreated := brs.dao.Create(&dto.BillKindModifyData{
					UserID:      userID,
					KindID:      newKindID,
					Name:        systemKind.Name,
					Description: systemKind.Description,
					OverKind:    kindID,
					UpKind:      systemKind.UpKind,
				})
				if isCreated != nil {
					return isCreated
				}
				return brs.dao.DeleteByUserIDAndKindID(userID, newKindID)
			}
			return err
		}
		if (overKind.DeletedAt != gorm.DeletedAt{}) {
			return constant.ErrBillKindNotExist
		}
		kindID2delete = overKind.KindID
	}
	if brs.dao.ExistByUserIDAndKindID(userID, kindID2delete) {
		return brs.dao.DeleteByUserIDAndKindID(userID, kindID2delete)
	}
	return constant.ErrBillKindNotExist
}
