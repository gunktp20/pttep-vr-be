package pttep_vr_db

import "gorm.io/gorm"

func (o *Operator) DB() *gorm.DB {
	return o.database.GetDB()
}
