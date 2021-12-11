package entity

import (
	"database/sql"
)

type Path struct {
	BaseEntity
	Path         sql.NullString `gorm:"column:Path"`
	Query        sql.NullString `gorm:"column:Query"`
	Fragment     sql.NullString `gorm:"column:Fragment"`
	BaseDomainId int64          `gorm:"column:BaseDomainId"`
	Counter      uint64         `gorm:"column:Counter"`
	BaseDomain   BaseDomain     `gorm:"foreignkey:BaseDomainId" `
}

func (Path) TableName() string {
	return "Pathes"
}
