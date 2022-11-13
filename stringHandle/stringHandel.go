package stringhandle

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 替换操作

func StringReplace(m []primitive.M, new string, old string) {
	//存储处理好后的数据
	var newM []primitive.M

	// 遍历第一遍，拿到集合中的文档
	for _, v := range m {
		// 定义一个map 存储处理后的结果
		myMap := make(map[string]interface{})
		//将_id赋值
		myMap["_id"] = v["_id"]
		// 遍历第二遍，拿到文档中的内容
		for key, vv := range v {
			// 跳过 _id
			if key == "_id" {
				continue
			}
			// 类型断言
			str, ok := vv.(string)
			// str是否是字符串
			if !ok {
				//不是字符串
				switch vv.(type) {
				case primitive.M:
					{
						fmt.Printf("%T\n", vv)
					}
				case primitive.A:
					{
						fmt.Printf("%T\n", vv)
					}
				default:
					{
						//存储没被类型断言的值
						myMap[key] = vv
					}
				}
			} else {
				// 将字符串替换
				s := strings.Replace(str, old, new, -1)
				// 处理好后的字符串赋值进myMap
				myMap[key] = s
			}
		}

		newM = append(newM, myMap)
	}

	fmt.Println(newM)

}
