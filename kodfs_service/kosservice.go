package kodfs_service

import (
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"github.com/guoqingpeng/kodfs/kodfs_http"
	"net/http"
)

func Start_Kodfs_Service(cfg *kodfs_config.KodfsConfig) {

	serverChan := make(chan int)

	protocals := make([]string, 2)
	protocals[0] = "http"
	protocals[1] = "https"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		kodfs_http.ProcessHandle(writer, request)
	})

	//分别开启http和https接口服务
	for i := 0; i < len(protocals); i++ {

		if i%2 == 0 {
			go func() {
				StartHttpListener(cfg)
			}()
		} else {
			go func() {
				StartHttpsListener(cfg)
			}()
		}

	}

	//启动后就一直阻塞住
	<-serverChan

}
