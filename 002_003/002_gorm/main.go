package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	// Code  string
	// gorm默认不操作零值，可以用sql.NullString允许零值，也可以用指针*string解决
	Code  sql.NullString `gorm:"unique;type:varchar(15)"`
	Price uint
}

var DB *gorm.DB

func main() {
	// GORM 是 Go 语言中最流行的 ORM（对象关系映射）框架，
	// 支持 MySQL、PostgreSQL、SQLite、SQL Server 等数据库。
	// 它提供了查询构建、事务支持、钩子、关联关系、自动迁移等功能。

	// 文档：https://gorm.io/zh_CN/docs/index.html

	// 安装（mysql）
	// go get -u gorm.io/gorm gorm.io/driver/mysql

	// 对于postgresql：
	// go get -u gorm.io/driver/postgres

	// 全局日志：
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // include params in the SQL log
			Colorful:                  true,        // Enable color
		},
	)

	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	DB = db
	log.Println("数据库连接成功")

	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: sql.NullString{"D42", false}, Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	// Update 会更新零值，但是Updates不会，可以定义为 sql.NullString 或 *string 解决
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})                    // 仅更新非零值字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}}) // 可更新为零值
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product - 如果表里有 deleted_at 字段，默认逻辑删除，只是修改了 deleted_at
	db.Delete(&product, 1)
	// 如果想物理删除，可以调用 Unscoped
	// db.Unscoped().Delete(&product)

}
