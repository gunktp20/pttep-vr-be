package models

const rolePermissions = "role_permissions"

type RolePermission struct {
	ID           uint        `gorm:"column:id" json:"id,omitempty" bson:"_id,omitempty"`
	RoleID       uint        `json:"role_id" bson:"role_id" gorm:"column:role_id"`
	PermissionID uint        `json:"permission_id" bson:"permission_id" gorm:"column:permission_id"`
	Role         *Role       `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID" bson:"-"`
	Permission   *Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID;references:ID" bson:"-"`
}

func (RolePermission) TableName() string {
	return rolePermissions
}

func (RolePermission) CollectionName() string {
	return rolePermissions
}

func (_self *RolePermission) SetID(v uint) {
	_self.ID = v
}

func (_self *RolePermission) SetRoleID(v uint) {
	_self.RoleID = v
}

func (_self *RolePermission) SetPermissionID(v uint) {
	_self.PermissionID = v
}
