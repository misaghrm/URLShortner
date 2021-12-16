package entity

import (
	"database/sql"
	"time"
)

type BaseDomain struct {
	Id                  int64        `gorm:"PRIMARY_KEY;UNIQUEINDEX;column:Id;not null;type:int8"` // PRIMARY_KEY;
	CreationTimeUtc     time.Time    `gorm:"column:CreationTimeUtc;not null"`
	ModificationTimeUtc sql.NullTime `gorm:"column:ModificationTimeUtc"`
	Url                 string       `gorm:"column:Url"`
	Counter             uint64       `gorm:"column:Counter"`
}

func (BaseDomain) TableName() string {
	return "URLs"
}
