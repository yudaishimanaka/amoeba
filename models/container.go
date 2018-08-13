package models

import (
	"time"
)

type Container struct {
	Id int64 `xorm:"not null BIGINT pk autoincr 'id'"`
	ImageId int64 `xorm:"image_id not null"`
	OsType string `xorm:"os_type not null"`
	Created time.Time `xorm:"created not null"`
	Ipv4 string `xorm:"ipv4_address"`
	Cpu int8 `xorm:"cpu not null"`
	Memory int32 `xorm:"memory not null"`
	Disk int64 `xorm:"disk not null"`
}
