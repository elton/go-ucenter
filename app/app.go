package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	redisServer       *redis.Client
	dbServer          *gorm.DB
	dbOnce, redisOnce sync.Once
)

// InitRedis 获得Redis实例
func initRedis() {
	redisServer = redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "", // no password set
		DB:                 0,  // use default DB
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	})
	// 检测心跳
	pong, err := redisServer.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		panic(err)
	}
	fmt.Printf("redis ping result: %s\n", pong)
}

// InitDb 初始化数据库
func initDb() {

	dsn := "root@tcp(127.0.0.1:3306)/ucenter?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	// 连接数设置
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	sqlDB.Ping()

	dbServer = db
}

// GetRedis 获得Redis实例
func GetRedis() *redis.Client {
	redisOnce.Do(initRedis)
	return redisServer
}

// GetDB 得到数据库实例
func GetDB() *gorm.DB {
	dbOnce.Do(initDb)
	return dbServer
}

// Destruct 销毁
func Destruct() {
	if redisServer != nil {
		_ = redisServer.Close()
	}

	if dbServer != nil {
		sqlDB, _ := dbServer.DB()
		_ = sqlDB.Close()
	}

}
