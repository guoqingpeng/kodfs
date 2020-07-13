package kodfs_http

import (
	"fmt"
	"net/http"
	"strings"
)

func ProcessHandle(writer http.ResponseWriter, request *http.Request) {

	uri := request.URL.RequestURI()
	fmt.Println(uri)
	if strings.Contains(uri, "/read/") {
		readFile(writer, request)
	} else if strings.Contains(uri, "/write/") {
		writeFile(writer, request)
	} else {
		writer.Write([]byte("暂不支持这个请求"))
	}

}

func readFile(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("文件读取成功"))

}

func writeFile(writer http.ResponseWriter, request *http.Request) {

	//返回具体的文件名以及要写到的dataserver中的信息
	writer.Write([]byte("文件保存成功"))

}
