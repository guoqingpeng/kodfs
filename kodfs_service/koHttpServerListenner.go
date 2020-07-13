package kodfs_service

import (
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"log"
	"net/http"
	"strconv"
)

func StartHttpListener(cfg *kodfs_config.KodfsConfig) {

	fmt.Println("http service start")
	addr := strconv.Itoa(cfg.HttpPort)
	err := http.ListenAndServe(":"+addr, nil)
	log.Fatal(err)

}
