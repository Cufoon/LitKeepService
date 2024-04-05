package vo

import (
	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/db/dto"
	"cufoon.litkeep.service/internal/db/entity"
)

type BillRecordCreateReq struct {
	KindID string   `json:"kindID"`
	Type   *int     `json:"type"`
	Value  *float64 `json:"value"`
	Time   *int64   `json:"time"`
	Mark   string   `json:"mark"`
}

type BillRecordCreateRes struct {
	Created bool `json:"created"`
}

type BillRecordQueryReq struct {
	KindID    string `json:"kindID"`
	StartTime *int64 `json:"startTime"`
	EndTime   *int64 `json:"endTime"`
}

type BillRecordQueryRes struct {
	Record []entity.BillRecord `json:"record"`
}

type BillRecordPageQueryReq struct {
	Page int `json:"page"`
}

type BillRecordPageQueryRes = BillRecordQueryRes

type BillRecordPageData struct {
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	PageSize   int64 `json:"pageSize"`
}

type BillRecordPageDataQueryRes struct {
	Kinds    []*box.BillKind    `json:"kinds"`
	PageData BillRecordPageData `json:"pageData"`
}

type BillRecordQueryWithKindRes struct {
	Kinds  []*box.BillKind     `json:"kinds"`
	Record []entity.BillRecord `json:"record"`
}

type BillRecordModifyReq struct {
	ID     uint     `json:"id"`
	KindID string   `json:"kindID"`
	Type   *int     `json:"type"`
	Value  *float64 `json:"value"`
	Time   *int64   `json:"time"`
	Mark   string   `json:"mark"`
}

type BillRecordModifyRes struct {
	Modified bool `json:"modified"`
}

type BillRecordDeleteReq struct {
	Ids []uint `json:"ids"`
}

type BillRecordDeleteRes struct {
	NotDeleted []uint `json:"notDeleted"`
}

type BillRecordStatisticsDayQueryReq struct {
	StartTime *int64 `json:"startTime"`
	EndTime   *int64 `json:"endTime"`
}

type BillRecordStatisticsDayQueryRes struct {
	Statistic []dto.QueryStatisticsDayData `json:"statistic"`
}
