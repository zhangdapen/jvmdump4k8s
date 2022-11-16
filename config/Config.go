package config

import (
	"github.com/go-ini/ini"
)

type Config struct {
	//七牛云
	QiniuApiHost   string `ini:"qiniu.apiHost"`
	QiniuAccessKey string `ini:"qiniu.accessKey"`
	QiniuSecretKey string `ini:"qiniu.secretKey"`
	QiniuBucket    string `ini:"qiniu.bucket"`
	QiniuFolder    string `ini:"qiniu.folder"`
	QiniuUseHTTPS  string `ini:"qiniu.usehttps"`
}

var inifile = "jvmdump4k8s.ini"
var GlobalConfig Config = Config{}

func init() {
	//加载INI文件
	cfg, err := ini.Load(inifile)
	err = cfg.Section("").MapTo(&GlobalConfig)
	if err != nil {
		panic(err)
	}
}
