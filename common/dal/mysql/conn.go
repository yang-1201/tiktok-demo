package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once = sync.Once{}
)

func NewDB() *gorm.DB {
	return db
}

func InitDB() {
<<<<<<< HEAD
	//dsn := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	dsn := "root:@tcp(127.0.0.1:3303)/tiktok?charset=utf8&parseTime=True&loc=Local&timeout=100s"
	//dsn := "root:@tcp(host.docker.internal:3303)/tiktok?charset=utf8&parseTime=True&loc=Local&timeout=100s"
=======
	dsn := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
>>>>>>> f3bcb08 (publish完成)
	//dsn := "root:yy6364650@tcp(localhost:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	InitDBWithDSN(dsn)
}

func InitDBWithDSN(dsn string) {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("连接数据库失败, err: " + err.Error())
		}
	})
}

func UpdateModel() {
	InitDB()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Follow{})
	db.AutoMigrate(&Like{})
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Comment{})
}
