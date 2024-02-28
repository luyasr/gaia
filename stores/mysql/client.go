package mysql

import (
	"github.com/luyasr/gaia/ioc"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Mysql).Client
}