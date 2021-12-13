package entity

import (
	"database/sql"
	"time"
)

type BaseDomain struct {
	Id                  int64        `gorm:"PRIMARY_KEY;UNIQUEINDEX;column:Id;not null;type:int8"` // PRIMARY_KEY;
	CreationTimeUtc     time.Time    `gorm:"column:CreationTimeUtc;not null"`
	ModificationTimeUtc sql.NullTime `gorm:"column:ModificationTimeUtc"`
	Scheme              string       `gorm:"column:TransferProtocol"`
	Opaque              string       `gorm:"column:Opaque"`
	Host                string       `gorm:"column:Host;not null"`
	TLD                 string       `gorm:"column:TLD"`
	Paths               []Path       `gorm:"foreignKey:BaseDomainId" `
}

func (BaseDomain) TableName() string {
	return "BaseDomains"
}
