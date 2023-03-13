package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type GormConfig struct {
	Host string
	User string
	PassWord string
	DbName string
	DeBug int
	MaxLifetime int
	MaxOpenConn int
	MaxIdleConn int
}

// Instance db实例
func Instance(gc GormConfig) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn(gc)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不使用复数
		},
	})

	if err != nil {
		panic(err)
	}
	if gc.DeBug == 1 {
		db.Debug()
	}
	ddb, _ := db.DB()
	ddb.SetMaxIdleConns(gc.MaxIdleConn)                                // 用于设置闲置的连接数
	ddb.SetMaxOpenConns(gc.MaxOpenConn)                                // 用于设置最大打开的连接数,默认值为0表示不限制
	ddb.SetConnMaxLifetime(time.Duration(gc.MaxLifetime) * time.Second) // 设置连接超时500秒

	return db
}

// dsn 链接
func dsn(gc GormConfig) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=", gc.User, gc.PassWord, gc.Host, gc.DbName) + "Asia%2FShanghai"
}