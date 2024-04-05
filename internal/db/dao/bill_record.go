package dao

import (
	"time"

	"gorm.io/gorm"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/db/entity"
)

func NewBillRecordDAO(db *gorm.DB, billRecordModel *entity.BillRecord) *BillRecordDAO {
	return &BillRecordDAO{db: db, billRecordModel: billRecordModel}
}

type BillRecordDAO struct {
	db              *gorm.DB
	billRecordModel *entity.BillRecord
}

func (br *BillRecordDAO) model() *gorm.DB {
	return br.db.Model(br.billRecordModel)
}

func (br *BillRecordDAO) Create(billRecord *box.BillRecord) error {
	result := br.model().Create(&entity.BillRecord{
		UserID: billRecord.UserID,
		Type:   billRecord.Type,
		Kind:   billRecord.KindID,
		Value:  billRecord.Value,
		Time:   billRecord.Time,
		Mark:   billRecord.Mark,
	})
	return result.Error
}

func (br *BillRecordDAO) QueryByUserID(userID string) []entity.BillRecord {
	var records []entity.BillRecord
	br.model().Where("id_user", userID).Find(&records)
	return records
}

func (br *BillRecordDAO) QueryByUserIDAndPage(userID string, page int) []entity.BillRecord {
	var records []entity.BillRecord
	br.model().Where("id_user", userID).Order("time desc").Offset(page * 20).Limit(20).Find(&records)
	return records
}

func (br *BillRecordDAO) QueryByUserIDForCount(userID string) int64 {
	var c int64
	br.model().Where("id_user", userID).Count(&c)
	return c
}

func (br *BillRecordDAO) QueryByUserIDPeriod(userID string, start, end *time.Time) []entity.BillRecord {
	var records []entity.BillRecord
	br.model().Where("id_user = ? AND (time BETWEEN ? AND ?)", userID, start, end).Order("time desc").Find(&records)
	return records
}

func (br *BillRecordDAO) QueryByUserIDAndKindID(userID string, kindID string) []entity.BillRecord {
	var records []entity.BillRecord
	br.model().Where("id_user", userID).Where("id_kind", kindID).Find(&records)
	return records
}

func (br *BillRecordDAO) QueryByUserIDAndKindPeriod(userID string, kindID string, start, end *time.Time) []entity.BillRecord {
	var records []entity.BillRecord
	br.model().Where(
		"id_user = ? AND id_kind = ? AND (time BETWEEN ? AND ?)",
		userID,
		kindID,
		start,
		end,
	).Find(&records)
	return records
}

func (br *BillRecordDAO) Delete(ids []uint, userID string) (notDeleted []uint) {
	var err error
	for _, item := range ids {
		err = br.model().Where("id", item).Where("id_user", userID).Delete(br.billRecordModel).Error
		if err != nil {
			notDeleted = append(notDeleted, item)
		}
	}
	return
}

func (br *BillRecordDAO) Modify(billRecord *box.BillRecord) error {
	return br.model().
		Where("id", billRecord.ID).
		Where("id_user", billRecord.UserID).Updates(&entity.BillRecord{
		Kind:  billRecord.KindID,
		Type:  billRecord.Type,
		Value: billRecord.Value,
		Time:  billRecord.Time,
		Mark:  billRecord.Mark,
	}).Error
}

func (br *BillRecordDAO) QueryStatisticsDay(userID string, start, end time.Time) []dto.QueryStatisticsDayData {
	var result []dto.QueryStatisticsDayData
	br.model().
		Select("date_format(time, '%Y-%m-%d') day,sum(if(type = 1, -1 * value, value)) total").
		Where("id_user = ? AND (time BETWEEN ? AND ?)", userID, start, end).
		Group("date_format(time, '%Y-%m-%d')").Order("day").
		Find(&result)
	return result
}
