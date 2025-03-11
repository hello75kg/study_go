package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kirinlabs/HttpRequest"
	"net/http"
	"strconv"
	"sync"
)

// 使用http的方式来调用远程方法
func httpAdd() {
	// http://127.0.0.1:8000/add?a=1&b=2
	// 1.call id : 用uri解决
	// 2.数据传输协议 : url参数传输
	// 3.网络协议 : http
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("path: ", r.URL.Path)
		a, err := strconv.Atoi(r.Form["a"][0])
		b, err := strconv.Atoi(r.Form["b"][0])

		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)

	})
	_ = http.ListenAndServe(":8080", nil)
}

func Add(a, b int) int {
	request := HttpRequest.NewRequest()
	url := "http://127.0.0.1:8080/add?a=" + strconv.Itoa(a) + "&b=" + strconv.Itoa(b)
	res, _ := request.Get(url)
	body, _ := res.Body()
	// fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

type ResponseData struct {
	Data int `json:"data"`
}

func main() {
	/*
		// rpc
		// (Remote Procedure Call)
		// 远程过程调用
		// 一个节点调用另外一个节点的服务通信协议
		// 可以像调用本地函数一样调用远程服务，隐藏了底层网络通信的细节
		// 跨语言

		// 1.call的id映射
		// 2.数据编码协议，序列化和反序列化- json/xml/protobuf/msgpack
		// 3.网络传输协议

		// http协议，建立在tcp上
		// 传统http1.x，一次性连接，返回后端口
		// http2.0支持长连接

		// 大部分RPC框架都是使用TCP协议实现
		// grpc基于http2.

			客户端：
				1.建立连接
				2.将数据序列化
				3.发送数据
				4.等待服务器响应，接收返回数据
				5.将数据反序列化
			服务端：
				1.监听端口
				2.接收到数据，读取数据
				3.将数据反序列化
				4.处理业务逻辑
				5.序列化处理结果
				6.返回给客户端

	*/

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpAdd()
	}()

	// 使用http的方式来调用远程方法
	for i := 0; i < 3; i++ {
		num := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := Add(num, num+1)
			fmt.Println(res)
		}()
	}
	wg.Wait()

	// rpc
	// 客户端
	// 客户端存根stub：封装地址，发送/接收/打包/解析数据
	// 服务端
	// 服务端存根

	// 序列化和反序列化
	// 动态代理：自动生成stub
}
