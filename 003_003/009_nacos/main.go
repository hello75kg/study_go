package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func main() {
	// nacos 安装：
	// docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
	//
	// 命名空间：隔离配置集合，一般用来区分微服务
	// 组：用来区分 开发环境、测试环境、生产环境
	// dataid: 配置集合，相当于一个配置文件
	//
	// 能获取到配置
	// 能坚挺到配置文件的修改
	// nacos即使一个配置中心，也是一个注册中心
	// 定义服务端配置（注意：请根据实际环境修改 IP 和端口）

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "192.168.0.249", // Nacos 服务地址
			Port:        8848,            // Nacos 服务端口（默认8848）
			ContextPath: "/nacos",        // 默认值
		},
	}

	// 定义客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "5360869b-57ee-4b6c-82db-395b64b6d9a9", // 如果有命名空间，请填写对应的 NamespaceId，否则留空
		TimeoutMs:           5000,                                   // 请求超时
		ListenInterval:      30000,                                  // 配置监听间隔
		NotLoadCacheAtStart: true,                                   // 启动时不加载本地缓存
		LogDir:              "./nacos/log",
		CacheDir:            "./nacos/cache",
		LogLevel:            "debug",
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Fatalf("创建 Nacos 配置客户端失败: %v", err)
	}

	// 定义 DataId 和 Group（建议 DataId 尽量唯一）
	dataId := "user-web.json"
	group := "dev"
	// content := "Hello, Nacos from Go SDK!"

	// 发布配置
	// success, err := configClient.PublishConfig(vo.ConfigParam{
	// 	DataId:  dataId,
	// 	Group:   group,
	// 	Content: content,
	// })
	// if err != nil {
	// 	log.Fatalf("发布配置失败: %v", err)
	// }
	// fmt.Printf("发布配置成功: %v\n", success)

	// 获取配置
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Fatalf("获取配置失败: %v", err)
	}
	fmt.Printf("获取到的配置内容: %s\n", config)
}
