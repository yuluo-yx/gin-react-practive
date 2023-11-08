package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// 初始化连接数据库配置

var _db *gorm.DB

func Database(conn string) {

	// 日志
	var ormLogger logger.Interface

	if gin.Mode() == "debug" {
		// 如果是debug模式 则 打印日志信息
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		// gorm 配置
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 设置打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	Migration()
}

// NewDBClient 创建一个数据库服务客户端
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
