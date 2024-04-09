package vo

import (
	"time"

	"cufoon.litkeep.service/internal/box"
	"cufoon.litkeep.service/internal/db/dto"
	"gorm.io/gorm"
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

type BillRecord struct {
	ID        uint           `json:"id"`
	UserID    string         `json:"userID"`
	Type      int            `json:"type"`
	Kind      string         `json:"kind"`
	Value     float64        `json:"value"`
	Time      time.Time      `json:"time"`
	Mark      string         `json:"mark"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type BillRecordQueryRes struct {
	Record []*BillRecord `json:"record"`
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
	Kinds    []*BillKind        `json:"kinds"`
	PageData BillRecordPageData `json:"pageData"`
}

type BillRecordQueryWithKindRes struct {
	Kinds  []*box.BillKind `json:"kinds"`
	Record []*BillRecord   `json:"record"`
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
	StartTime  *int64 `json:"startTime"`
	EndTime    *int64 `json:"endTime"`
	RecordType *int   `json:"type"`
}

type BillRecordStatisticsDayQueryRes struct {
	Statistic []dto.QueryStatisticsDayData `json:"statistic"`
}
