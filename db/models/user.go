package models

import (
	"goweb/utils"
	"time"

	"gorm.io/gorm"
)

// User [...]
type User struct {
	UserId    string    `json:"user_id" gorm:"primary_key;unique;column:user_id;type:varchar(20);not null" binding:"required,alphanum,lt=20,gte=5" label:"账户Id"` // 用户Id
	Password  string    `gorm:"column:password;type:varchar(40);not null" binding:"required,alphanum,lt=20,gte=8" label:"密码"`                                    // 密码
	Phone     string    `gorm:"column:phone;type:char(11)" binding:"required,checkPhone" label:"手机号码"`                                                           // 手机
	Balance   int       `gorm:"column:balance;type:int(10);not null" label:"账户余额"`                                                                               // 余额
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:null;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:null;not null"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = utils.Md5(user.Password)
	return nil
}

type RUser struct {
	UserID  string `json:"user_id"`
	Phone   string `json:"phone"`
	Balance int    `json:"balance"`
	Token   string `json:"token"`
}

type UserLogin struct {
	UserId   string `json:"user_id" gorm:"primary_key;unique;column:user_id;type:varchar(20);not null" binding:"required,alphanum,lt=20,gte=5" label:"账户Id"` // 用户Id
	Password string `gorm:"column:password;type:varchar(40);not null" binding:"required,alphanum,lt=20,gt=8" label:"密码"`
}

type UserRecharge struct {
	UserID string `json:"user_id" binding:"required,alphanum,lt=20,gte=5" label:"账户Id"` // 用户Id`
	Amount int    `json:"amount" binding:"required,numeric,min=1" label:"充值金额"`
}
