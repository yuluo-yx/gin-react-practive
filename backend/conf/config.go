package conf

import (
	"backend/dao"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

// 定义全局变量
var (
	AppModel string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// Init 初始化配置
func Init() {
	// 本地读取环境变量  注意：此路径是从main.go开始算起的路径
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		// 处理读取配置文件异常
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	//读取配置
	LoadServer(file)
	LoadMySql(file)

	dsn := strings.Join([]string{DbUser, ":", DbPassword,
		"@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	fmt.Printf("执行")
	dao.Database(dsn)
}

// LoadMySql 加载数据库配置信息
func LoadMySql(file *ini.File) {
	Db = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

// LoadServer 加载服务配置信息
func LoadServer(file *ini.File) {
	// 将读取进来的配置字符串化
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
