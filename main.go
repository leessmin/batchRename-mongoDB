package main

import (
	db "batchRename/DB"
	readfile "batchRename/readFile"
	update_DB "batchRename/update-DB"
	"fmt"
	"os"
)

func main() {
	// 读取配置
	err := readfile.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 初始化 连接mongo
	db.InitMongo()

	// 更新文章
	update_DB.UpdateDB()
	fmt.Println("操作的文章个数为:", update_DB.Count)

	// 退出程序
	fmt.Println("操作完成!!!")
	fmt.Println("按任意键退出...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
