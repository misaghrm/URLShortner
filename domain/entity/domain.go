package entity

type BaseDomain struct {
	BaseEntity
	Scheme string `gorm:"column:TransferProtocol"`
	Opaque string `gorm:"column:Opaque"`
	Host   string `gorm:"column:Host;not null"`
	TLD    string `gorm:"column:TLD"`
	Path   []Path `gorm:"foreignkey:BaseDomainId" `
}

func (BaseDomain) TableName() string {
	return "BaseDomains"
}
