package handler

import (
	"github.com/gofiber/fiber/v2"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/service"
	"cufoon.litkeep.service/internal/util"
	"cufoon.litkeep.service/internal/vo"
	"cufoon.litkeep.service/pkg/flow"
)

func NewBillKindHandler(billKindService *service.BillKindService) *BillKindHandler {
	return &BillKindHandler{billKindService: billKindService}
}

type BillKindHandler struct {
	billKindService *service.BillKindService
}

func (bkh *BillKindHandler) Create(c *fiber.Ctx) error {
	data := new(vo.BillKindCreateReq)
	err := c.BodyParser(data)
	if err != nil || data.Name == "" {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	err = bkh.billKindService.Create(&box.BillKindCreate{
		UserID:      userID,
		Name:        data.Name,
		Description: data.Description,
		UpKind:      data.UpKind,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillKindCreateRes{Created: true})
}

func (bkh *BillKindHandler) Query(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	kinds, err := bkh.billKindService.Query(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	rKinds := make([]vo.BillKindQueryTreeResData, len(kinds))
	for idx, item := range kinds {
		tmpChildren := make([]vo.BillKind, len(item.Children))
		for jdx, item2 := range item.Children {
			tmpChildren[jdx] = vo.BillKind{
				KindID:      item2.KindID,
				Name:        item2.Name,
				Description: item2.Description,
			}
		}
		rKinds[idx] = vo.BillKindQueryTreeResData{
			BillKind: vo.BillKind{
				KindID:      item.KindID,
				Name:        item.Name,
				Description: item.Description,
			},
			Children: tmpChildren,
		}
	}
	return util.ResOK(c, &vo.BillKindQueryTreeRes{Kind: rKinds})
}

func (bkh *BillKindHandler) Modify(c *fiber.Ctx) error {
	data := new(vo.BillKindModifyReq)
	err := c.BodyParser(data)
	if err != nil || data.KindID == "" || data.Name == "" {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	err = bkh.billKindService.Modify(userID, &box.BillKind{
		KindID:      data.KindID,
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillKindModifyRes{Modified: true})
}

func (bkh *BillKindHandler) Delete(c *fiber.Ctx) error {
	data := new(vo.BillKindDeleteReq)
	err := c.BodyParser(data)
	if err != nil || data.KindID == "" {
		return util.ResBadBody(c)
	}
	userID := flow.GetUserID(c)
	err = bkh.billKindService.Delete(userID, data.KindID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &vo.BillKindDeleteRes{Deleted: true})
}
