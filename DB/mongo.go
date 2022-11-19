package db

// 连接mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 数据库的 配置
type MongoDBConfig struct {
	// mongo地址
	MongoUrl string
	// 数据库名字
	DBName   string
	// 集合
	CollName []string
	// 旧值
	OldValue string
	// 新值
	NewValue string
}

// 数据库的配置DBConfig
var DBConfig MongoDBConfig

// mongo 的连接
var Client *mongo.Client

func InitMongo() {
	// 设置时长
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	var err error
	// 连接mongo
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(DBConfig.MongoUrl))
	if err != nil {
		panic(fmt.Sprintln("连接mongoDB时出现了错误:", err))
	}
}
