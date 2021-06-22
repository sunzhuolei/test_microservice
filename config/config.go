package config

import (
	"sync"

	//"closer_user/internal/pkg/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Config struct {
	Server Server
	Data Data
}

type Server struct {
	Http Http
	Grpc Grpc
}

type Data struct {
	DataBase DataBase
	Redis Redis
}
type Http struct {
	Addr string `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}


type Grpc struct {
	Addr string `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
}


type DataBase struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}


type Redis struct {
	Addr string `yaml:"addr"`
	DB int `yaml:"db"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

var (
	config   *Config
	configLock = new(sync.RWMutex)
)

/**
载入配置信息
 */
func LoadConfig()bool{
	temp := new(Config)
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil{
		fmt.Println("读取配置文件失败：",err.Error())
		return false
	}
	err = yaml.Unmarshal(yamlFile,temp)
	if err != nil{
		fmt.Println("解析配置文件失败：",err.Error())
		return false
	}
	fmt.Println(temp)
	configLock.Lock()
	config = temp
	configLock.Unlock()
	return true
}



func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}