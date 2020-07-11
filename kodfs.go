package main

import (
	"flag"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
)

//set up system args only accept from the console

var (
	configRoot *string = flag.String("c", "./conf", "provide the config file path")
	logRoot    *string = flag.String("l", "./log", "provide the log file path")
)

func main() {

	fmt.Println("可牛了。。。分布式对象存储系统开始服务")

	//step1 start	to parse the args from inputs
	flag.Parse() //别忘记了写这个
	//end parser

	//step2 解析主配置文件
	confPath := *configRoot + "/kodfs.cfg"

	cfg := kodfs_config.NewKodsConfig()

	cfg.ParseConfig(confPath)

	fmt.Println(cfg)

}
