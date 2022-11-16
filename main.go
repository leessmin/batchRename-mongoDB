package main

import (
	update_DB "batchRename/update-DB"
	"fmt"
)

func main() {
	// 躲开一条协程
	update_DB.UpdateDB()

	fmt.Println("操作的文章个数为:", update_DB.Count)
}
