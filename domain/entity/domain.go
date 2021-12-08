package entity

import (
	"database/sql"
	"github.com/misaghrm/urlshortener/domain/enum/protocoltype"
)

type BaseDomain struct {
	BaseEntity
	TransferProtocol protocoltype.ProtocolType `gorm:"column:TransferProtocol;not null"`
	SubDomain        sql.NullString            `gorm:"column:SubDomain"`
	Domain           string                    `gorm:"column:Domain;not null"`
	TLD              string                    `gorm:"column:TLD;not null"`
}

func (BaseDomain) TableName() string {
	return "BaseDomains"
}
