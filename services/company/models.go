package company

import (
	"apus-sample/common/database/models"
)

type Company struct {
	models.BaseModel
	Title        string  `gorm:"column:title;not null" json:"title"`
	Address      *string `gorm:"column:address" json:"address"`
	ContactPhone *string `gorm:"column:contact_phone" json:"contact_phone"`
	ContactEmail *string `gorm:"column:contact_email" json:"contact_email"`
}

func (c Company) FilterMap() map[string]string {
	return map[string]string{
		"id":            "id",
		"title":         "title",
		"address":       "address",
		"contact-phone": "contact_phone",
	}
}

//func (Company) TableName() string {
//	return "companies"
//}
