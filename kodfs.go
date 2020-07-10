package main

import (
	"fmt"

	"github.com/guoqingpeng/kodfs/kodfs_config"
	)


func main() {

	fmt.Print("分布式对象存储系统开始服务")
	fmt.Println()
	fmt.Println(kodfs_config.Config())

}