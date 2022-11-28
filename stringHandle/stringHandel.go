package stringhandle

import (
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 替换操作

func StringReplace(m []primitive.M, new string, old string) []primitive.M {
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
				case primitive.M,
					primitive.A:
					{
						filter, t := MRecursion(vv, old, new)
						// MRecursion(vv, old, new)

						myMap[key] = filterMyMap(t, filter)
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

	return newM
}

// MRecursion primitive.M的递归处理函数  返回值：处理过的数据，还有是否是map值 1是map 0是slice  如果是 slice 需要用 key: leeSlice 获取slice
func MRecursion(m interface{}, old string, new string) (map[string]interface{}, int) {

	//定义返回的类型
	var flag int = 1
	//定义一个map用来存储处理好的值
	myMap := make(map[string]interface{})

	// 使用反射 判断 m 是否是 map 类型
	if reflect.TypeOf(m).Kind() == reflect.Map {
		// 取值
		value := reflect.ValueOf(m)

		// 判断是否是空值
		if !value.IsValid() {
			fmt.Printf("这个地方是一个nil:%v\n", value.IsNil())
		}

		// 迭代值
		val := value.MapRange()

		//开始遍历值
		for val.Next() {
			//将 key 转 string 除去两边空格
			k := strings.TrimSpace(fmt.Sprintln(val.Key()))
			// 判断值是否为string
			if val.Value().Elem().Kind() == reflect.String {
				// 将 val 转 string  替换字符串
				v := strings.Replace(fmt.Sprintln(val.Value()), old, new, -1)
				//存储值 除去两边空格
				myMap[k] = strings.TrimSpace(v)
			} else {

				// 判断类型 根据类型 断言后赋值
				switch val.Value().Elem().Kind() {
				case reflect.Int:
					myMap[k] = val.Value().Interface().(int)
				case reflect.Int8:
					myMap[k] = val.Value().Interface().(int8)
				case reflect.Int16:
					myMap[k] = val.Value().Interface().(int16)
				case reflect.Int32:
					myMap[k] = val.Value().Interface().(int32)
				case reflect.Int64:
					// 判断是否是primitive.DateTime 类型
					switch val.Value().Interface().(type) {
					case primitive.DateTime:
						myMap[k] = val.Value().Interface().(primitive.DateTime)
					default:
						myMap[k] = val.Value().Interface().(int64)
					}
				case reflect.Bool:
					myMap[k] = val.Value().Interface().(bool)
				case reflect.Interface:
					myMap[k] = val.Value().Interface()
				case reflect.Map:
					myMap[k] = traverseMap(val.Value().Interface(), old, new)
				case reflect.Slice:
					myMap[k] = traverseSlice(val.Value().Interface(), old, new)
				default:
					// TODO 这应该是一个坑，将 case没有匹配的类型暂时变成字符串
					myMap[k] = fmt.Sprintln(val.Value().Interface())
				}
			}
		}
	}

	// 使用反射 判断 m 是否是 slice 类型
	if reflect.TypeOf(m).Kind() == reflect.Slice {
		// slice 类型
		flag = 0
		//定义一个切片 存储值
		mySlice := make([]interface{}, 0)

		// 取值
		value := reflect.ValueOf(m)

		// 判断是否是空值
		if !value.IsValid() {
			fmt.Printf("这个地方是一个nil:%v\n", value.IsNil())
		}

		// 开始遍历slice
		for i := 0; i < value.Len(); i++ {
			// 判断值是否为string
			if value.Index(i).Elem().Kind() == reflect.String {
				// 将 val 转 string  替换字符串
				v := strings.Replace(fmt.Sprintln(value.Index(i)), old, new, -1)
				//存储值 除去两边空格
				mySlice = append(mySlice, strings.TrimSpace(v))
			} else {

				// 判断类型 根据类型 断言后赋值
				switch value.Index(i).Elem().Kind() {
				case reflect.Int:
					mySlice = append(mySlice, value.Index(i).Interface().(int))
				case reflect.Int8:
					mySlice = append(mySlice, value.Index(i).Interface().(int8))
				case reflect.Int16:
					mySlice = append(mySlice, value.Index(i).Interface().(int16))
				case reflect.Int32:
					mySlice = append(mySlice, value.Index(i).Interface().(int32))
				case reflect.Int64:
					// 判断是否是primitive.DateTime 类型
					switch value.Index(i).Interface().(type) {
					case primitive.DateTime:
						mySlice = append(mySlice, value.Index(i).Interface().(primitive.DateTime))
					default:
						mySlice = append(mySlice, value.Index(i).Interface().(int64))
					}
				case reflect.Bool:
					mySlice = append(mySlice, value.Index(i).Interface().(bool))
				case reflect.Interface:
					mySlice = append(mySlice, value.Index(i).Interface())
				case reflect.Map:
					// 遍历map 替换里面的值 然后再赋值
					mySlice = append(mySlice, traverseMap(value.Index(i).Interface(), old, new))
				case reflect.Slice:
					// 遍历slice 替换里面的值 然后再赋值
					mySlice = append(mySlice, traverseSlice(value.Index(i).Interface(), old, new))
				default:
					// TODO 这应该是一个坑，将 case没有匹配的类型暂时变成字符串
					mySlice[i] = fmt.Sprintln(value.Index(i).Interface())
				}

			}
		}

		//将过滤好的slice存储到myMap
		myMap["leeSlice"] = mySlice

	}

	return myMap, flag
}

// 过滤好数据，到底是map还是slice等
func filterMyMap(i int, myMap map[string]interface{}) interface{} {
	if i == 1 {
		//存储过滤后的值
		return myMap
	} else {
		return myMap["leeSlice"]
	}
}

// 遍历 嵌套的数据 map类型
func traverseMap(m interface{}, old, new string) map[string]interface{} {

	//定义一个map用来存储处理好的值
	myMap := make(map[string]interface{})

	// 取值
	value := reflect.ValueOf(m)

	// 判断是否是空值
	if !value.IsValid() {
		fmt.Printf("这个地方是一个nil:%v\n", value.IsNil())
	}

	// 判断是否为 map 类型
	if value.Kind() == reflect.Map {

		// 迭代值
		val := value.MapRange()

		// 遍历值
		for val.Next() {
			//将 key 转 string 除去两边空格
			k := strings.TrimSpace(fmt.Sprintln(val.Key()))

			// 判断类型
			if val.Value().Elem().Kind() == reflect.String {

				// 将 val 转 string  替换字符串
				v := strings.Replace(fmt.Sprintln(val.Value()), old, new, -1)
				//存储值 除去两边空格
				myMap[k] = strings.TrimSpace(v)
			} else {
				// 不是字符串类型
				// 判断类型 根据类型 断言后赋值
				switch val.Value().Elem().Kind() {
				case reflect.Int:
					myMap[k] = val.Value().Interface().(int)
				case reflect.Int8:
					myMap[k] = val.Value().Interface().(int8)
				case reflect.Int16:
					myMap[k] = val.Value().Interface().(int16)
				case reflect.Int32:
					myMap[k] = val.Value().Interface().(int32)
				case reflect.Int64:
					// 判断是否是primitive.DateTime 类型
					switch val.Value().Interface().(type) {
					case primitive.DateTime:
						myMap[k] = val.Value().Interface().(primitive.DateTime)
					default:
						myMap[k] = val.Value().Interface().(int64)
					}
				case reflect.Bool:
					myMap[k] = val.Value().Interface().(bool)
				case reflect.Interface:
					myMap[k] = val.Value().Interface()
				case reflect.Map:
					myMap[k] = traverseMap(val.Value().Interface(), old, new)
				case reflect.Slice:
					myMap[k] = traverseSlice(val.Value().Interface(), old, new)
				default:
					// TODO 这应该是一个坑，将 case没有匹配的类型暂时变成字符串
					myMap[k] = fmt.Sprintln(val.Value().Interface())
				}
			}

		}

	}

	return myMap
}

// 遍历 嵌套的数据 slice类型
func traverseSlice(m interface{}, old, new string) []interface{} {

	//定义一个切片 存储值
	mySlice := make([]interface{}, 0)

	// 取值
	value := reflect.ValueOf(m)

	// 判断是否是空值
	if !value.IsValid() {
		fmt.Printf("这个地方是一个nil:%v\n", value.IsNil())
	}

	// 开始遍历slice
	for i := 0; i < value.Len(); i++ {
		// 判断值是否为string
		if value.Index(i).Elem().Kind() == reflect.String {
			// 将 val 转 string  替换字符串
			v := strings.Replace(fmt.Sprintln(value.Index(i)), old, new, -1)
			//存储值 除去两边空格
			mySlice = append(mySlice, strings.TrimSpace(v))
		} else {

			// 判断类型 根据类型 断言后赋值
			switch value.Index(i).Elem().Kind() {
			case reflect.Int:
				mySlice = append(mySlice, value.Index(i).Interface().(int))
			case reflect.Int8:
				mySlice = append(mySlice, value.Index(i).Interface().(int8))
			case reflect.Int16:
				mySlice = append(mySlice, value.Index(i).Interface().(int16))
			case reflect.Int32:
				mySlice = append(mySlice, value.Index(i).Interface().(int32))
			case reflect.Int64:
				// 判断是否是primitive.DateTime 类型
				switch value.Index(i).Interface().(type) {
				case primitive.DateTime:
					mySlice = append(mySlice, value.Index(i).Interface().(primitive.DateTime))
				default:
					mySlice = append(mySlice, value.Index(i).Interface().(int64))
				}
			case reflect.Bool:
				mySlice = append(mySlice, value.Index(i).Interface().(bool))
			case reflect.Interface:
				mySlice = append(mySlice, value.Index(i).Interface())
			case reflect.Map:
				// 遍历map 替换里面的值 然后再赋值
				mySlice = append(mySlice, traverseMap(value.Index(i).Interface(), old, new))
			case reflect.Slice:
				// 遍历slice 替换里面的值 然后再赋值
				mySlice = append(mySlice, traverseSlice(value.Index(i).Interface(), old, new))
			default:
				// TODO 这应该是一个坑，将 case没有匹配的类型暂时变成字符串
				mySlice[i] = fmt.Sprintln(value.Index(i).Interface())
			}

		}
	}

	return mySlice
}
