package handler

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/handler/validator"
	"cufoon.litkeep.service/internal/service"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/internal/vo"
	"cufoon.litkeep.service/pkg/flow"
)

func NewBillRecordHandler(billRecordService *service.BillRecordService, billKindService *service.BillKindService) *BillRecordHandler {
	return &BillRecordHandler{billRecordService: billRecordService, billKindService: billKindService}
}

type BillRecordHandler struct {
	billRecordService *service.BillRecordService
	billKindService   *service.BillKindService
}

func (brh *BillRecordHandler) Create(c *fiber.Ctx) error {
	data := new(vo.BillRecordCreateReq)
	err := c.BodyParser(data)
	if err != nil || data.Time == nil || data.Value == nil || data.Type == nil || validator.BadBillRecordType(data.Type) {
		return util.ResBadBody(c)
	}

	userID := flow.GetUserID(c)
	err = brh.billRecordService.Create(&box.BillRecord{
		UserID: userID,
		KindID: data.KindID,
		Type:   *data.Type,
		Value:  *data.Value,
		Time:   *util.Timestamp2Time(data.Time),
		Mark:   data.Mark,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordCreateRes{Created: true})
}

func (brh *BillRecordHandler) Query(c *fiber.Ctx) error {
	data := new(vo.BillRecordQueryReq)
	err := c.BodyParser(data)
	if err != nil || validator.BadTimeDuration(data.StartTime, data.EndTime) {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.Query(&box.BillRecordQuery{
		UserID:    userID,
		KindID:    data.KindID,
		StartTime: util.Timestamp2Time(data.StartTime),
		EndTime:   util.Timestamp2Time(data.EndTime),
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordQueryRes{Record: record})
}

func (brh *BillRecordHandler) QueryPage(c *fiber.Ctx) error {
	data := new(vo.BillRecordPageQueryReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.QueryPage(&box.BillRecordPageQuery{
		UserID: userID,
		Page:   data.Page,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordPageQueryRes{Record: record})
}

func (brh *BillRecordHandler) QueryPageData(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	count, err := brh.billRecordService.QueryRecordCount(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	kinds, err := brh.billKindService.QueryLinear(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordPageDataQueryRes{
		Kinds: kinds,
		PageData: vo.BillRecordPageData{
			Total:      count,
			TotalPages: (count-1)/20 + 1,
			PageSize:   20,
		}})
}

func (brh *BillRecordHandler) QueryWithKind(c *fiber.Ctx) error {
	data := new(vo.BillRecordQueryReq)
	err := c.BodyParser(data)
	if err != nil || validator.BadTimeDuration(data.StartTime, data.EndTime) {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.Query(&box.BillRecordQuery{
		UserID:    userID,
		KindID:    data.KindID,
		StartTime: util.Timestamp2Time(data.StartTime),
		EndTime:   util.Timestamp2Time(data.EndTime),
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	kinds, err := brh.billKindService.QueryLinear(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordQueryWithKindRes{Kinds: kinds, Record: record})
}

func (brh *BillRecordHandler) Modify(c *fiber.Ctx) error {
	data := new(vo.BillRecordModifyReq)
	err := c.BodyParser(data)
	if err != nil || data.ID == 0 || data.Type == nil || data.Time == nil || validator.BadBillRecordType(data.Type) {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	err = brh.billRecordService.Modify(&box.BillRecord{
		ID:     data.ID,
		UserID: userID,
		KindID: data.KindID,
		Type:   *data.Type,
		Value:  *data.Value,
		Time:   *util.Timestamp2Time(data.Time),
		Mark:   data.Mark,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordModifyRes{Modified: true})
}

func (brh *BillRecordHandler) Delete(c *fiber.Ctx) error {
	data := new(vo.BillRecordDeleteReq)
	err := c.BodyParser(data)
	if err != nil || data.Ids == nil || len(data.Ids) == 0 {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	notDeleted := brh.billRecordService.Delete(data.Ids, userID)
	return util.ResOK(c, &vo.BillRecordDeleteRes{NotDeleted: notDeleted})
}

func (brh *BillRecordHandler) QueryStatisticsDay(c *fiber.Ctx) error {
	data := new(vo.BillRecordStatisticsDayQueryReq)
	err := c.BodyParser(data)
	if err != nil || validator.BadTimeDuration(data.StartTime, data.EndTime) {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	result, err := brh.billRecordService.QueryStatisticsDay(userID, &box.BillRecordStatisticsDayQueryReq{
		StartTime: util.Timestamp2Time(data.StartTime),
		EndTime:   util.Timestamp2Time(data.EndTime),
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillRecordStatisticsDayQueryRes{Statistic: result})
}
