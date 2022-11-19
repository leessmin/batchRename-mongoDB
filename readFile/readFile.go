package readfile

import (
	db "batchRename/DB"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 读取 config.conf 文件ReadFile
func ReadFile() error {
	// 打开文件
	file, err := os.Open("config.conf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 关闭文件
	defer file.Close()

	// 扫描
	scanner := bufio.NewScanner(file)
	// 行读取
	scanner.Split(bufio.ScanLines)

	// 扫描文本
	for scanner.Scan() {
		// 判断符合那个文本

		if strings.Contains(scanner.Text(), "CollName") {
			// 分割字符串
			arr := strings.Split(scanner.Text(), "=")

			// 删除字符[]
			var text string
			text = strings.ReplaceAll(arr[1], "[", "")
			text = strings.ReplaceAll(text, "]", "")
			// 分割字符串 拿到每个数据库的名字
			arrText := strings.Split(text, ",")

			// 遍历arrText
			for _, v := range arrText {
				// 除去两边空格 存入 DBConfig.DBName中
				db.DBConfig.CollName = append(db.DBConfig.CollName, strings.TrimSpace(v))
			}
		}

		if strings.Contains(scanner.Text(), "MongoUrl") {
			// 将处理好的值放入配置结构体中
			db.DBConfig.MongoUrl = mySplit(scanner.Text(), "=")
		}

		if strings.Contains(scanner.Text(), "OldVal") {
			// 将处理好的值放入配置结构体中
			db.DBConfig.OldValue = mySplit(scanner.Text(), "=")
		}

		if strings.Contains(scanner.Text(), "NewVal") {
			// 将处理好的值放入配置结构体中
			db.DBConfig.NewValue = mySplit(scanner.Text(), "=")
		}

		if strings.Contains(scanner.Text(), "DBName") {
			// 将处理好的值放入配置结构体中
			db.DBConfig.DBName = mySplit(scanner.Text(), "=")
		}
	}

	if db.DBConfig.MongoUrl == "" || len(db.DBConfig.CollName) == 0 || db.DBConfig.OldValue == "" || db.DBConfig.NewValue == "" || db.DBConfig.DBName == "" {
		panic("读取配置文件出错")
	}
	fmt.Println("读取配置成功")

	return nil
}

// 处理字符串
func mySplit(str, cut string) string {
	// 分割字符串
	arr := strings.Split(str, cut)

	return strings.TrimSpace(arr[1])
}
