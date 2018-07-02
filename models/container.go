package models

import (
	"time"
)

type Container struct {
	Id int64 `xorm:"not null BIGINT pk autoincr 'id'"`
	ImageId int64 `xorm:"image_id"`
	Created time.Time `xorm:"created"`
	Ipv4 string `xorm:"ipv4_address"`
	Cpu int8 `xorm:"cpu"`
	Memory int32 `xorm:"memory"`
	Disk int64 `xorm:"disk"`
}