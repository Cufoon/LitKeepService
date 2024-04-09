package dao

import (
	"gorm.io/gorm"

	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/db/entity"
)

func NewBillKindDAO(db *gorm.DB, billKindModel *entity.BillKind) *BillKindDAO {
	return &BillKindDAO{db: db, billKindModel: billKindModel}
}

type BillKindDAO struct {
	db            *gorm.DB
	billKindModel *entity.BillKind
}

func (bk *BillKindDAO) model() *gorm.DB {
	return bk.db.Model(bk.billKindModel)
}

func (bk *BillKindDAO) Create(billKind *dto.BillKindModifyData) error {
	return bk.model().Create(&entity.BillKind{
		KindID:      billKind.KindID,
		UserID:      billKind.UserID,
		Name:        billKind.Name,
		Description: billKind.Description,
		UpKind:      billKind.UpKind,
		OverKind:    billKind.OverKind,
	}).Error
}

func (bk *BillKindDAO) ExistByKindID(kindID string) bool {
	record := new(entity.BillKind)
	r := bk.model().Where("id_kind", kindID).Order("name asc").First(record)
	return r.Error == nil
}

func (bk *BillKindDAO) ExistByUserIDAndKindID(userID string, kindID string) bool {
	record := new(entity.BillKind)
	r := bk.model().Where("id_user", userID).Where("id_kind", kindID).Order("name asc").First(record)
	return r.Error == nil
}

func (bk *BillKindDAO) ExistByUserIDAndKindIDWithSystem(userID string, kindID string) bool {
	record := new(entity.BillKind)
	r := bk.model().Where("id_kind", kindID).Where("id_user = ? or id_user = ?", userID, "LitAdmin").Order("name asc").First(record)
	return r.Error == nil
}

func (bk *BillKindDAO) QueryByKindID(kindID string) (*entity.BillKind, error) {
	record := new(entity.BillKind)
	r := bk.model().Where("id_kind", kindID).Order("name asc").First(record)
	return record, r.Error
}

func (bk *BillKindDAO) QueryByUserID(userID string) ([]entity.BillKind, error) {
	var records []entity.BillKind
	r := bk.model().Where("id_user", userID).Order("name asc").Find(&records)
	return records, r.Error
}

func (bk *BillKindDAO) QueryByUserIDWithDeletedOver(userID string) ([]entity.BillKind, error) {
	var records []entity.BillKind
	// r := bk.model().Raw("select * from bill_kind where id_user = ? and (deleted_at is not null or id_overkind is not null or id_overkind !='');", userID).Find(&records)
	r := bk.model().Unscoped().Where("id_user = ? and (id_overkind is not null or id_overkind !='')", userID).Find(&records)
	return records, r.Error
}

func (bk *BillKindDAO) QueryByOverKindIDAndUserIDContainDeleted(userID, kindID string) (*entity.BillKind, error) {
	record := new(entity.BillKind)
	r := bk.model().Unscoped().Where("id_user", userID).Where("id_overkind", kindID).Order("name asc").First(record)
	return record, r.Error
}

func (bk *BillKindDAO) Modify(billKind *dto.BillKindModifyData) error {
	result := bk.model().Where("id_user", billKind.UserID).
		Where("id_kind", billKind.KindID).Updates(&entity.BillKind{
		Name:        billKind.Name,
		Description: billKind.Description,
		UpKind:      billKind.UpKind,
		OverKind:    billKind.OverKind,
	})
	return result.Error
}

func (bk *BillKindDAO) DeleteByUserID(userID string) error {
	result := bk.model().Where("id_user", userID).Delete(bk.billKindModel)
	return result.Error
}

func (bk *BillKindDAO) DeleteByUserIDAndKindID(userID, kindID string) error {
	result := bk.model().Where("id_user", userID).Where("id_kind", kindID).Delete(bk.billKindModel)
	return result.Error
}
