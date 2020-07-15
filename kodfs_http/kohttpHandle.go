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

	//根据文件id，返回的该的chunk信息
	//以及每个chunk所在的dataserver信息
	//以及每个chunk所在的block信息
	//然后由客户端自己去取chunk信息合成一个完整的file
	writer.Write([]byte("文件读取成功"))

}

func writeFile(writer http.ResponseWriter, request *http.Request) {

	//返回具体的文件名以及要写到的dataserver中的信息
	writer.Write([]byte("文件保存成功"))

}
