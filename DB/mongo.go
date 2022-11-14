package db

// 连接mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo 的连接
var Client *mongo.Client

func init() {
	fmt.Println("开始连接mongoDB中...")
	// 设置时长
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	var err error
	// 连接mongo
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.1.104:27017"))
	if err != nil {
		panic(fmt.Sprintln("连接mongoDB时出现了错误:", err))
	}
	fmt.Println("连接mongoDB成功!!!")
}
