package service

import (
	"math"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/constant"
	"cufoon.litkeep.service/internal/db/dao"
	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/db/entity"
)

func NewBillRecordService(dao *dao.BillRecordDAO, billKindDAO *dao.BillKindDAO) *BillRecordService {
	return &BillRecordService{dao: dao, billKindDAO: billKindDAO}
}

type BillRecordService struct {
	dao         *dao.BillRecordDAO
	billKindDAO *dao.BillKindDAO
}

func (brs *BillRecordService) Create(record *box.BillRecord) error {
	if record.KindID != "" {
		if !brs.billKindDAO.ExistByUserIDAndKindID(record.UserID, record.KindID) {
			return constant.ErrBillKindNotExist
		}
	}
	return brs.dao.Create(record)
}

func (brs *BillRecordService) QueryRecordCount(userID string) (int64, error) {
	count := brs.dao.QueryByUserIDForCount(userID)
	return count, nil
}

func (brs *BillRecordService) Query(query *box.BillRecordQuery) ([]entity.BillRecord, error) {
	which := 0
	if query.KindID != "" {
		which++
	}
	if query.StartTime != nil &&
		query.EndTime != nil &&
		query.StartTime.Before(*query.EndTime) {
		which += 2
	}
	var result []entity.BillRecord
	switch which {
	case 0:
		result = brs.dao.QueryByUserID(query.UserID)
	case 1:
		result = brs.dao.QueryByUserIDAndKindID(query.UserID, query.KindID)
	case 2:
		result = brs.dao.QueryByUserIDPeriod(query.UserID, query.StartTime, query.EndTime)
	case 3:
		result = brs.dao.QueryByUserIDAndKindPeriod(query.UserID, query.KindID, query.StartTime, query.EndTime)
	}
	return result, nil
}

func (brs *BillRecordService) QueryPage(query *box.BillRecordPageQuery) ([]entity.BillRecord, error) {
	result := brs.dao.QueryByUserIDAndPage(query.UserID, query.Page)
	return result, nil
}

func (brs *BillRecordService) QueryStatisticsDay(userID string, query *box.BillRecordStatisticsDayQueryReq) ([]dto.QueryStatisticsDayData, error) {
	oResult := brs.dao.QueryStatisticsDay(userID, *query.StartTime, *query.EndTime)
	diff := query.EndTime.Sub(*query.StartTime)
	diffDays := int64(math.Floor(diff.Abs().Hours() / 24))
	dateList := make([]string, 0, diffDays)
	for d := query.StartTime.AddDate(0, 0, 1); !d.After(*query.EndTime); d = d.AddDate(0, 0, 1) {
		dateList = append(dateList, d.Format("2006-01-02"))
	}
	result := make([]dto.QueryStatisticsDayData, 0, diffDays)
	idx1 := 0
	idx2 := 0
	for idx1 < len(dateList) && idx2 < len(oResult) {
		if dateList[idx1] == oResult[idx2].Day {
			result = append(result, dto.QueryStatisticsDayData{Day: oResult[idx2].Day, Money: oResult[idx2].Money})
			idx1++
			idx2++
		} else {
			result = append(result, dto.QueryStatisticsDayData{Day: dateList[idx1], Money: 0})
			idx1++
		}
	}
	return result, nil
}

func (brs *BillRecordService) Modify(billRecord *box.BillRecord) error {
	if billRecord.KindID != "" {
		if !brs.billKindDAO.ExistByUserIDAndKindID(billRecord.UserID, billRecord.KindID) {
			return constant.ErrBillKindNotExist
		}
	}
	return brs.dao.Modify(billRecord)
}

func (brs *BillRecordService) Delete(ids []uint, userID string) (notDeleted []uint) {
	notDeleted = brs.dao.Delete(ids, userID)
	return
}
