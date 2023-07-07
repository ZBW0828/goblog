package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

//定义了一个配置文件的结构和全局变量 Cfg，并提供了初始化函数来读取和解析配置文件。通过这些配置信息，应用程序可以根据需要进行自定义设置。
type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "go-log"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
