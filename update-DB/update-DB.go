package update_DB

import (
	db "batchRename/DB"
	stringhandle "batchRename/stringHandle"
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// 更新mongoDB里的数据

var wg sync.WaitGroup

var mu sync.Mutex

// Count 计数操作文档的个数
var Count int = 0

// UpdateDB 更新数据库数据
func UpdateDB() {
	// 从配置文件中读取数据库信息
	myArr := db.DBConfig.CollName

	//遍历 数据库
	for _, v := range myArr {
		// 添加WaitGroup
		wg.Add(1)
		// 开启协程
		go func(v string) {
			// 完成释放WaitGroup
			defer wg.Done()
			// 获取数据库
			m, err := db.GetData(db.DBConfig.DBName, v)
			// 判断是否出错
			if err != nil {
				fmt.Println(err)
				return
			}

			// 处理 替换 需要的字符
			ms := stringhandle.StringReplace(m, db.DBConfig.NewValue, db.DBConfig.OldValue)

			// 将处理好后的数据更新至mongoDB
			for _, mv := range ms {
				// 获取id
				id := mv["_id"]
				// 过滤条件
				filter := bson.D{{Key: "_id", Value: id}}
				// 更新的数据
				update := bson.D{{Key: "$set", Value: mv}}
				// 连接集合
				coll := db.Client.Database("test").Collection(v)
				// 更新
				result, err := coll.UpdateOne(context.TODO(), filter, update)
				// 判断是否出错
				if err != nil {
					panic(err)
				}

				// 发送修改的文章数量到通道中
				var cou int = int(result.ModifiedCount)

				// 打印修改的文章
				// fmt.Println("修改的文章：", cou)

				mu.Lock()
				Count += cou
				mu.Unlock()
			}
		}(v)
	}

	wg.Wait()
}
