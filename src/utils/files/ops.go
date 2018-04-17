package files

import (
	"os"
	"fmt"
)

var SysAPI Sys_opt_FileInfo


type Sys_opt_FileInfo struct{


}


func (this *Sys_opt_FileInfo) GetfileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if  err ==nil{
		//文件大小
		filesize:= fileInfo.Size()
		return filesize
	}else {
		fmt.Println(err.Error())
	}
	return 0

}