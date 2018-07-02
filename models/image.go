package models

type Image struct {
	Id int64 `xorm:"not null BIGINT pk autoincr 'id'"`
	Name string `xorm:"image_name not null"`
}
