package entity

import (
	"database/sql"
)

type Path struct {
	BaseEntity
	RawPath      sql.NullString `gorm:"column:RawPath"`
	RawQuery     sql.NullString `gorm:"column:RawQuery"`
	RawFragment  sql.NullString `gorm:"column:RawFragment"`
	BaseDomainId int64          `gorm:"column:BaseDomainId"`
	BaseDomain   BaseDomain     `gorm:"foreignkey:BaseDomainId" `
}

func (Path) TableName() string {
	return "Pathes"
}
