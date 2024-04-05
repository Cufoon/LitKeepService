package cache

import (
	"gorm.io/gorm"

	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/entity"
)

type SystemBillKindCacheObject = map[string]entity.BillKind

type SystemBillKindCache struct {
	cache SystemBillKindCacheObject
	dao   *dao.BillKindDAO
}

func NewSystemBillKindCache(dao *dao.BillKindDAO) *SystemBillKindCache {
	return &SystemBillKindCache{dao: dao}
}

func (sbc *SystemBillKindCache) InitCache() bool {
	sbc.cache = make(SystemBillKindCacheObject)
	kinds, err := sbc.dao.QueryByUserID("LitAdmin")
	if err != nil {
		return false
	}
	for _, item := range kinds {
		if item.KindID != "KindLitWorldRoot" {
			sbc.cache[item.KindID] = item
		}
	}
	return true
}

func (sbc *SystemBillKindCache) GetCache() SystemBillKindCacheObject {
	return sbc.cache
}

func (sbc *SystemBillKindCache) GetByKindID(kindID string) entity.BillKind {
	return sbc.cache[kindID]
}

func (sbc *SystemBillKindCache) ExistByKindID(kindID string) bool {
	_, ok := sbc.cache[kindID]
	return ok
}

func (sbc *SystemBillKindCache) JoinSystemBillKind(kinds []entity.BillKind) []entity.BillKind {
	result := make([]entity.BillKind, 0)
	var isOveredMap = make(map[string]bool)
	var isDeletedMap = make(map[string]bool)
	for _, kind := range kinds {
		if kind.OverKind != "" {
			if _, ok := sbc.cache[kind.OverKind]; ok {
				isOveredMap[kind.OverKind] = true
				if (kind.DeletedAt != gorm.DeletedAt{}) {
					isDeletedMap[kind.KindID] = true
				}
			}
		}
	}
	for _, kind := range sbc.cache {
		if !isOveredMap[kind.KindID] && !isDeletedMap[kind.KindID] {
			result = append(result, kind)
		}
	}
	result = append(result, kinds...)
	var filteredKinds []entity.BillKind
	for _, item := range result {
		if (item.DeletedAt == gorm.DeletedAt{}) {
			filteredKinds = append(filteredKinds, item)
		}
	}
	return filteredKinds
}
