package models

import (
	"api/pkg/logging"
	"api/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //数据库实例，定义为全局变量

func SetUp() {
	var (
		err          error
		databaseType = setting.DatabaseSetting.Type
		user         = setting.DatabaseSetting.User
		pass         = setting.DatabaseSetting.Password
		host         = setting.DatabaseSetting.Host
		name         = setting.DatabaseSetting.Name
	)

	//使用gorm 连接数据库
	db, err = gorm.Open(databaseType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, pass, host, name))
	if err != nil {
		logging.Fatal("数据库连接失败", err)
	}
	//设置表名称得前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true) //设置金庸标名得附属形式
	db.LogMode(true)       //打印日志，本地调试的时候可以打开看执行的sql语句

	db.DB().SetMaxIdleConns(10)  //设置空闲时的最大连接数
	db.DB().SetMaxOpenConns(100) // 设置数据库的最大打开连接数
}
