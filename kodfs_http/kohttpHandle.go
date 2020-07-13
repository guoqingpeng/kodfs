package kodfs_http

import (
	"fmt"
	"net/http"
	"strings"
)

func ProcessHandle(writer http.ResponseWriter, request *http.Request) {

	uri := request.URL.RequestURI()
	fmt.Println(uri)
	if strings.Contains(uri, "read") {
		readFile(writer, request)
	} else if strings.Contains(uri, "write") {
		writeFile(writer, request)
	} else {

	}

}

func readFile(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("文件读取成功"))

}

func writeFile(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("文件保存成功"))

}
