package repository

import (
	"pttep-vr-api/pkg/repository/pttep-vr-db"
	"pttep-vr-api/pkg/utils/gormDB"
)

type Interface interface {
	PTTEPVR() *pttep_vr_db.Operator
}
type Operator struct {
	pttep_vr *pttep_vr_db.Operator
}

func New(db gormDB.Interface) *Operator {
	return &Operator{
		pttep_vr: pttep_vr_db.New(db),
	}
}

func (r *Operator) PTTEPVR() *pttep_vr_db.Operator {
	return r.pttep_vr
}
