package entity

type BaseDomain struct {
	BaseEntity
	Scheme  string `gorm:"column:TransferProtocol;not null"`
	Opaque  string `gorm:"column:Opaque"`
	Host    string `gorm:"column:Host;not null"`
	TLD     string `gorm:"column:TLD;not null"`
	Counter uint64 `gorm:"column:Counter"`
	Path    []Path `gorm:"foreignkey:BaseDomainId" `
}

func (BaseDomain) TableName() string {
	return "BaseDomains"
}
