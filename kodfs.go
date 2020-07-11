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

	//step1 start	to parse the args from command line
	flag.Parse()
	//end step1 parser from command line

	//step2 parse the main config file  kodfs.cfg
	confPath := *configRoot + "/kodfs.cfg"
	cfg := kodfs_config.NewKodsConfig()
	cfg.ParseConfig(confPath)
	fmt.Println(cfg.HttpPort)
	fmt.Println(cfg.HttpsPort)
	fmt.Println(cfg.MaxCups)
	//end step2

}
