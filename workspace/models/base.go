package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	ORM *gorm.DB
	err error
)

func Init(c db.Connection) {
	ORM, err = gorm.Open("mysql", c.GetDB("antiddos"))

	if err != nil {
		panic("initialize orm failed")
	}
}
