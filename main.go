package main

import (
	db "batchRename/DB"
	stringhandle "batchRename/stringHandle"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

var wg sync.WaitGroup

func main() {
	myArr := []string{"blog_info", "article"}

	for _, v := range myArr {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			m, err := db.GetData("test", v)
			if err != nil {
				fmt.Println(err)
				return
			}
			ms := stringhandle.StringReplace(m, "http://1.1.1.1:5500", "http://47.115.219.17:5500")
			for _, mv := range ms {
				id := mv["_id"]
				filter := bson.D{{"_id", id}}
				update := bson.D{{"$set", mv}}
				coll := db.Client.Database("test").Collection(v)
				result, err := coll.UpdateOne(context.TODO(), filter, update)
				if err != nil {
					panic(err)
				}
				fmt.Println(result)
			}
			fmt.Println("-------------------------------------------")
		}(v)
	}

	wg.Wait()
}
