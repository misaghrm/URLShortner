package entity

import (
	"database/sql"
	"time"
)

type Path struct {
	Id                  int64          `gorm:"PRIMARY_KEY;UNIQUEINDEX;column:Id;not null;type:int8"` // PRIMARY_KEY;
	CreationTimeUtc     time.Time      `gorm:"column:CreationTimeUtc;not null"`
	ModificationTimeUtc sql.NullTime   `gorm:"column:ModificationTimeUtc"`
	Path                sql.NullString `gorm:"column:Path"`
	Query               sql.NullString `gorm:"column:Query"`
	Fragment            sql.NullString `gorm:"column:Fragment"`
	BaseDomainId        int64          `gorm:"column:BaseDomainId"`
	Counter             uint64         `gorm:"column:Counter"`
	BaseDomain          BaseDomain
}

func (Path) TableName() string {
	return "Pathes"
}
