package models

import "time"

const userRoles = "user_roles"

type UserRole struct {
	ID          uint      `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name" gorm:"column:name"`
	Email       string    `json:"email" bson:"email" gorm:"column:email"`
	UserID      uint      `json:"user_id" bson:"user_id" gorm:"column:user_id"`
	RoleID      uint      `json:"role_id" bson:"role_id" gorm:"column:role_id"`
	IsActive    bool      `json:"is_active" bson:"is_active" gorm:"column:is_active"`
	IsRemove    bool      `json:"is_remove" bson:"is_remove" gorm:"column:is_remove"`
	CreatedDate time.Time `json:"created_date" bson:"created_date" gorm:"column:created_date"`
	UpdatedDate time.Time `json:"updated_date" bson:"updated_date" gorm:"column:updated_date"`
	User        *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID" bson:"-"`
	Role        *Role     `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID" bson:"-"`
}

func (UserRole) TableName() string {
	return userRoles
}

func (UserRole) CollectionName() string {
	return userRoles
}

func (_self *UserRole) SetID(v uint) {
	_self.ID = v
}

func (_self *UserRole) SetUserID(v uint) {
	_self.UserID = v
}

func (_self *UserRole) SetRoleID(v uint) {
	_self.RoleID = v
}
