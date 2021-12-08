package entity

import (
	"database/sql"
	"time"
)

type BaseEntity struct {
	Id                  int64        `gorm:"UNIQUEINDEX;AUTO_INCREMENT:false;column:Id;not null;type:int8"` // PRIMARY_KEY;
	CreationTimeUtc     time.Time    `gorm:"column:CreationTimeUtc;not null"`
	ModificationTimeUtc sql.NullTime `gorm:"column:ModificationTimeUtc"`
}
