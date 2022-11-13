package main

import (
	db "batchRename/DB"
	stringhandle "batchRename/stringHandle"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	myArr := []string{"blog_info"}

	for _, v := range myArr {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			m, err := db.GetData("test", v)
			if err != nil {
				fmt.Println(err)
				return
			}
			stringhandle.StringReplace(m, "http://1.1.1.1:5500", "http://47.115.219.17:5500")
			fmt.Println("-------------------------------------------")
		}(v)
	}

	wg.Wait()
}
