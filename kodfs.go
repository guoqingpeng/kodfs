package main

import (
	"flag"
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"github.com/guoqingpeng/kodfs/kodfs_service"
)

//set up system args only accept from the console

var (
	configRoot *string = flag.String("c", "./conf/", "provide the config file path")
	logRoot    *string = flag.String("l", "./log/", "provide the log file path")
)

func main() {

	//step1 start	to parse the args from command line
	flag.Parse()
	//end step1 parser from command line

	//step2 parse the main config file  kodfs.cfg
	confPath := *configRoot + "kodfs.cfg"
	cfg := kodfs_config.NewKodsConfig()
	cfg.ParseConfig(confPath)
	//end step2

	//step3 	启动http以及https 开始服务了
	kodfs_service.Start_Kodfs_Service(cfg)
	//step3 end

}
