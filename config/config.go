package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// AppLocalPath 项目部署根目录
const AppLocalPath = "/home/workspace/myBlog/myBlogWeb"

type Config struct {
	Viewer Viewer       `yaml:"viewer"`
	System SystemConfig `yaml:"system"`
	Sql    SqlConfig    `yaml:"mysql"`
}
type Viewer struct {
	Title       string   `yaml:"Title"`
	Description string   `yaml:"Description"`
	Logo        string   `yaml:"Logo"`
	Navigation  []string `yaml:"Navigation"`
	CSDN        string   `yaml:"CSDN"`
	GitHub      string   `yaml:"GitHub"`
	Avatar      string   `yaml:"Avatar"`
	UserName    string   `yaml:"UserName"`
	UserDesc    string   `yaml:"UserDesc"`
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string `yaml:"CdnURL"`
	QiniuAccessKey  string `yaml:"QiniuAccessKey"`
	QiniuSecretKey  string `yaml:"QiniuSecretKey"`
	Valine          bool   `yaml:"Valine"`
	ValineAppid     string `yaml:"ValineAppid"`
	ValineAppkey    string `yaml:"ValineAppkey"`
	ValineServerURL string `yaml:"ValineServerURL"`
}

type SqlConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

var config *Config

func init() {
	config = new(Config)
	// 打开 YAML 文件
	file, err := os.Open(filepath.Join(AppLocalPath, "config/config.yaml"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}
	config.System.AppName = "Aaj-go-blog"
	config.System.Version = 1.0
	config.System.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	//log.Println(config)
}

func GetConfig() *Config {
	return config
}
