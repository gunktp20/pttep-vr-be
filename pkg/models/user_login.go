package models

import "time"

const userLogins = "user_logins"

type UserLogin struct {
	ID          uint `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`
	UserID      uint `json:"user_id" bson:"user_id" gorm:"column:user_id"`
	LoginTypeID uint `json:"login_type_id" bson:"login_type_id" gorm:"column:login_type_id"`

	Username    string    `gorm:"column:username" json:"username" bson:"username"`
	Password    string    `gorm:"column:password" json:"password" bson:"password"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
}

func (UserLogin) TableName() string {
	return userLogins
}

func (UserLogin) CollectionName() string {
	return userLogins
}

func (_self *UserLogin) SetID(v uint) {
	_self.ID = v
}

func (_self *UserLogin) SetUserID(v uint) {
	_self.UserID = v
}

func (_self *UserLogin) SetLoginTypeID(v uint) {
	_self.LoginTypeID = v
}

func (_self *UserLogin) SetUsername(v string) {
	_self.Username = v
}

func (_self *UserLogin) SetPassword(v string) {
	_self.Password = v
}
