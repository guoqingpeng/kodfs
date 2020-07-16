package kodfs_service

import (
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"github.com/guoqingpeng/kodfs/kodfs_http"
	"github.com/guoqingpeng/kodfs/kodfs_nameserver"
	"net/http"
)

func Start_Kodfs_Service(cfg *kodfs_config.KodfsConfig) {

	serverChan := make(chan int)

	protocals := make([]string, 2)
	protocals[0] = "http"
	protocals[1] = "https"

	//启动名称节点服务
	nameserver := kodfs_nameserver.NewNameServer()

	//主动监听来自dataserver发来的消息
	go func() {
		nameserver.NameServer_Start(cfg)
	}()

	//拦截客户端所有请求进行处理 提供文件外部使用的接口
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		kodfs_http.ProcessHandle(writer, request)
	})

	http.Handle("/read/", http.StripPrefix("/read", http.FileServer(http.Dir("./files"))))

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
