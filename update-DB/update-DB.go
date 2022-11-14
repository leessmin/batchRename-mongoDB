package update_DB

import (
	db "batchRename/DB"
	stringhandle "batchRename/stringHandle"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// 更新mongoDB里的数据

var wg sync.WaitGroup

// UpdateDB 更新数据库数据
func UpdateDB() {
	myArr := []string{"blog_info", "article"}

	//遍历 数据库
	for _, v := range myArr {
		// 添加WaitGroup
		wg.Add(1)
		// 开启协程
		go func(v string) {
			// 完成释放WaitGroup
			defer wg.Done()
			// 获取数据库
			m, err := db.GetData("test", v)
			// 判断是否出错
			if err != nil {
				fmt.Println(err)
				return
			}

			// 处理 替换 需要的字符
			ms := stringhandle.StringReplace(m, "http://1.1.1.1:5500", "http://47.115.219.17:5500")

			// 将处理好后的数据添加至mongoDB
			for _, mv := range ms {
				// 获取id
				id := mv["_id"]
				// 过滤条件
				filter := bson.D{{"_id", id}}
				// 更新的数据
				update := bson.D{{"$set", mv}}
				// 连接集合
				coll := db.Client.Database("test").Collection(v)
				// 更新
				result, err := coll.UpdateOne(context.TODO(), filter, update)
				// 判断是否出错
				if err != nil {
					panic(err)
				}
				// 打印修改的文章
				fmt.Println("修改的文章：", result.ModifiedCount)
			}
		}(v)
	}

	wg.Wait()
}
