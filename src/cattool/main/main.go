package main

import (
	"os"
	"fmt"
	"strings"

	"net/http"
	"time"
	sys "utils/files"
)

func readContentFromHttp(url string) []byte {
	respd,err :=http.Get(url)
	if err!=nil {
		panic(err)
	}
	buf:=make([]byte,100000)
	nread,err:=respd.Body.Read(buf)
		if err==nil && nread>0{
			return buf
		}
	return buf
}

func readContentFromFile(path string) []byte {
	file,err:= os.Open(path)
	defer file.Close()

	if err==nil{

		fsize:=sys.SysAPI.GetfileSize(path)
		buf:=make([]byte,fsize)
		fmt.Println("--> 文件大小:%d",fsize)
		if err==nil{
			nread,err:=file.Read(buf)
			if err==nil && nread>0{
				return buf
			}
		}

	}else {
		panic(err)
	}
    return nil
}

func main() {
	args:= os.Args
    fmt.Println(args)
    arglen:=len(args)
    var arg1 string
    if arglen>0{
		for i:=1;i<arglen; i++ {
			arg1 =args[i]
			var bytes []byte
			if strings.HasPrefix(arg1,"http://") {
				time.Sleep(time.Duration(time.Second*1))
				bytes=readContentFromHttp(arg1)

			}else {
				time.Sleep(time.Duration(time.Second*1))
				bytes = readContentFromFile(arg1)

			}

			fmt.Println("--> 内容如下:\n",string(bytes))


		}

	}




}
