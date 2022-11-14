package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 获取mongoDB集合中的数据

// GetData 传入 数据库的名字 集合的名字 返回查询到的结果
func GetData(dataName string, collName string) ([]primitive.M, error) {
	// 连接集合
	collection := Client.Database(dataName).Collection(collName)
	// 查询
	c, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("查询时出现了错误：", err)
		return nil, err
	}

	// 准备一个切片 存储好数据
	var arr []primitive.M
	// 遍历游标获取数据
	for c.Next(context.TODO()) {
		var result bson.M
		if err := c.Decode(&result); err != nil {
			log.Fatal(err)
			return nil, err
		}

		// 追加到arr
		arr = append(arr, result)
	}

	return arr, nil
}
