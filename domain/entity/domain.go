package entity

import (
	"github.com/misaghrm/urlshortener/domain/enum/protocoltype"
)

type BaseDomain struct {
	BaseEntity
	TransferProtocol protocoltype.ProtocolType `gorm:"column:TransferProtocol;not null"`
}

func (BaseDomain) TableName() string {
	return "BaseDomains"
}
