package entity

import "gorm.io/gorm"

type BillKind struct {
	gorm.Model
	KindID      string `gorm:"column:id_kind"`
	UserID      string `gorm:"column:id_user"`
	Name        string
	Description string `gorm:"default:null"`
	UpKind      string `gorm:"column:id_upkind;default:null"`
	OverKind    string `gorm:"column:id_overkind;default:null"`
}

func (BillKind) TableName() string {
	return "bill_kind"
}
