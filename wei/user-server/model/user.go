package model

import (
	__ "github.com/yuhang-jieke/yuedemo/wei/user-server/handler/proto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30);comment:姓名"`
	Age     int    `gorm:"type:int(11);comment:年龄"`
	Address string `gorm:"type:varchar(30);comment:地址"`
}

func (u *User) Registers(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Update(db *gorm.DB, in *__.UpdateReq) error {
	return db.Model(&u).Where("id=?", in.Id).Update("address", in.Address).Error
}

func (u *User) FindName(db *gorm.DB, name string) error {
	return db.Where("name=?", name).Find(&u).Error
}
