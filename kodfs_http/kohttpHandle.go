package kodfs_http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ProcessHandle(writer http.ResponseWriter, request *http.Request) {

	uri := request.URL.RequestURI()
	fmt.Println(uri)
	if strings.Contains(uri, "/readfile") {
		readFile(writer, request)
	} else if strings.Contains(uri, "/writefile") {
		writeFile(writer, request)
	} else {
		myupoadForm :=
			"<html>" +
				"<form action='writefile' method='post' enctype='multipart/form-data'>" +
				"<input type='file' name='myfile'/>" +
				"<input type='submit' value='upload'/>" +
				"</form>" +
				"</html>"
		writer.Write([]byte(myupoadForm))
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

	//获取文件的元数据

	//返回文件的存储信息

	//让客户端按照存储信息去和具体的数据服务器进行联系

}

func saveFileChunkIntoBlock() {
	filebytes, err := ioutil.ReadAll(nil)
	if err != nil {
		fmt.Println(err)
	}
	filenewname := strconv.Itoa(time.Now().Nanosecond())

	newfile, err := os.Create("./files/" + filenewname)
	if err != nil {
		fmt.Println(err)
	}
	newfile.Write(filebytes)
}
