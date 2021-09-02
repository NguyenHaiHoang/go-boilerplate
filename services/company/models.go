package company

import (
	"apus-sample/common/database/models"
)

type Company struct {
	models.BaseModel
	Code         string  `gorm:"column:code;not null" json:"code"`
	Title        string  `gorm:"column:title;not null" json:"title"`
	Address      *string `gorm:"column:address" json:"address"`
	ContactPhone *string `gorm:"column:contact_phone" json:"contact_phone"`
	ContactEmail *string `gorm:"column:contact_email" json:"contact_email"`
}

//func (Company) TableName() string {
//	return "companies"
//}