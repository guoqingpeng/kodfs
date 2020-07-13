package kodfs_service

import (
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"log"
	"net/http"
	"strconv"
)

func StartHttpsListener(cfg *kodfs_config.KodfsConfig) {

	fmt.Println("https service start")
	addr := strconv.Itoa(cfg.HttpsPort)
	err := http.ListenAndServeTLS(":"+addr, "./conf/server.crt", "./conf/server.key", nil)
	log.Fatal(err)

}
