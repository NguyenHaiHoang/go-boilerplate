package company

import (
	"apus-sample/common/database/models"
)

type User struct {
	models.BaseModel
	CompanyID    uint64
	Code         string
	Username     string
	FirstName    *string
	LastName     *string
	ContactPhone *string
	ContactEmail *string
}

func (User) TableName() string {
	return "users"
}
