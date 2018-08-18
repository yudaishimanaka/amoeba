package models

import (
	"time"
)

type Container struct {
	Id int `xorm:"not null BIGINT pk autoincr 'id'"`
	ImageId int `xorm:"image_id not null"`
	OsType string `xorm:"os_type not null"`
	Created time.Time `xorm:"created not null"`
	Ipv4 string `xorm:"ipv4_address"`
	Cpu int `xorm:"cpu not null"`
	Memory int `xorm:"memory not null"`
	Disk int `xorm:"disk not null"`
}
