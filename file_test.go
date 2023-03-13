package go_utils

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetDirInfo(t *testing.T) {

	// 用法一：获取当前目录内的所有文件、目录
	a,err:=GetDirInfo("./", func(fullPath string,info os.FileInfo) (pass bool) {
		if len(info.Name()) >=1 && info.Name()[0] == '.' {
			fmt.Println("过滤路径：",fullPath)
			return false
		}
		return true
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, info := range a {
		fmt.Println(*info)
	}
	fmt.Println("-------------------------------------------------------------")
	// 用法二：获取当前目录内的所有文件、目录，并且过滤掉git相关的文件
	a,err=GetDirInfo("./", func(fullPath string,info os.FileInfo) (pass bool) {
		if len(info.Name()) >= 4 && info.Name()[:4] == ".git" {
			return false
		}
		return true
	})

	for _, info := range a {
		fmt.Println(info.Name)
	}

	fmt.Println(err, len(a))

}
