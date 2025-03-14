package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func getEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func main() {
	// viper
	// 配置文件管理库
	// viper 是 Go 语言中功能强大的 配置管理库，支持 JSON、YAML、TOML、ENV 变量、命令行参数 等多种配置方式，并且支持 动态热加载。
	// 安装
	// go get -u github.com/spf13/viper
	configFileName := "config-pro"
	if env := getEnvInfo("WSHOP_ENV"); env != "pro" {
		configFileName = "config-dev"
	}

	viper.SetConfigName(configFileName)                                              // 不要加后缀
	viper.SetConfigType("yaml")                                                      // 指定配置文件类型
	viper.AddConfigPath("/Users/wwj/GolandProjects/studyProject/003_000/002_viper/") // 指定查找配置文件的路径（当前目录）

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 获取单个值
	fmt.Println("App Name:", viper.GetString("app_name"))
	fmt.Println("Port:", viper.GetInt("port"))
	fmt.Println("Debug Mode:", viper.GetBool("debug"))

	// 读取嵌套值
	fmt.Println("Database Host:", viper.GetString("database.host"))
	fmt.Println("Database Port:", viper.GetInt("database.port"))

	// 绑定环境变量（优先级：环境变量 > 配置文件）
	// 运行时设置环境变量
	// export DATABASE_PASSWORD="super_secret"
	// go run main.go
	viper.BindEnv("database.password") // 绑定环境变量
	fmt.Println("Database Password:", viper.GetString("database.password"))

	// 监听配置文件变化（热更新）
	// 如果配置文件发生变化，可以自动加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变更:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("读取配置文件失败: %v", err)
		}
		// 加载到struct ...
	})

	// 绑定命令行参数
	// go run main.go --port=8081
	pflag.Int("port", 9090, "Server port")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Println("Port from CLI:", viper.GetInt("port"))

	// 使用 Struct 解析配置
	// 可以自动解析到结构体：
	type Config struct {
		AppName  string `mapstructure:"app_name"`
		Port     int    `mapstructure:"port"`
		Debug    bool   `mapstructure:"debug"`
		Database struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"database"`
	}

	var config Config
	viper.Unmarshal(&config)
	fmt.Printf("%+v\n", config)
}
