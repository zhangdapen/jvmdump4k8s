package main

import (
	"flag"
	"fmt"
	"jvmdump4k8s/qiniu"
	"jvmdump4k8s/util"
)

var (
	dumpFile string //dump的文件
)

func main() {
	//parseCli()
	//fmt.Println("start invoke dump...")
	//flag.Parse()
	dumpFile := "./zhw.dump"
	fmt.Printf("dumpFile %s \n", dumpFile)
	//dump文件是否存在
	exist, err := util.FileExists(dumpFile)
	if err != nil {
		fmt.Printf("验证文件是否存在发生错误![%v]\n", err)
		return
	}
	if exist {
		var url = uploadStorage(dumpFile)
		fmt.Printf("OSS上传完成 %s\n", url)
	} else {
		fmt.Printf("dump文件不存在 %s\n", dumpFile)
	}
}

//解析命令行参数
func parseCli() {
	flag.StringVar(&dumpFile, "f", "", "-f xx.dump")
}

//Storage
func uploadStorage(file string) string {
	return qiniu.Upload(file)
}
