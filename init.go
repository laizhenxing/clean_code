package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	ccLog "cc/application/log"
	"cc/infrastructure/conf"
	"cc/infrastructure/persistence"
	"cc/infrastructure/persistence/orm"
	"cc/infrastructure/router"
	"cc/infrastructure/utils/cache"
)

func Start() {
	// 初始化配置信息
	conf.TomlConfig("config.toml", ".")

	// 初始化 Redis
	if err := cache.InitRedis(viper.GetString("redis.address"), viper.GetString("redis.password")); err != nil {
		ccLog.Fatalf("Redis 初始化失败 %v", err)
	}

	if _, err := persistence.NewMongoClient(viper.GetString("mongo.uri")); err != nil {
		ccLog.Fatalf("MongoDB Init Error %v", err)
	}

	// 使用 gorm
	conn, err := orm.NewGorm("mysql", viper.GetString("db.dsn"))
	if err != nil {
		ccLog.Fatalf("Mysql Init Error %v", err)
	}
	defer conn.Close()
	// 数据库迁移
	// conn.AutoMigrate()

	r := router.InitGinRouter()

	if err := http.ListenAndServe(viper.GetString("http.port"), r); err != nil {
		ccLog.Fatalf("Http Server Run Error %v", err)
	}
}
