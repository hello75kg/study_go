package main

func main() {
	// docker
	//
	// mysql安装：
	// docker pull mysql:8.0
	// docker run -p 3306:3306
	// 		--name mysql_server
	// 		-v $PWD/.mysql/conf:/etc/mysql/conf.d
	// 		-v $PWD/.mysql/logs:/logs
	// 		-v $PWD/.mysql/data:/var/lib/mysql
	// 		-e MYSQL_ROOT_PASSWORD=123456
	// 		-d mysql:8.0
	//
	// 查看mysql容器进程
	// docker ps -a
	// 进入mysql容器
	// docker exec -it ac52e7d915d9 /bin/bash
	//
	// 进入后连接mysql
	// mysql -u root -p123456
	// root@'%' 代表允许所有 IP 远程访问。
	// mysql_native_password 认证方式兼容老版本5.x访问（8.0+是 caching_sha2_password）
	// ALTER USER 'root'@'%' IDENTIFIED WITH caching_sha2_password BY '123456';
	// 授权
	// GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
	// FLUSH PRIVILEGES;
	//
	// 监听0.0.0.0
	// SHOW VARIABLES LIKE 'bind_address';
	// 如果返回的是 127.0.0.1，需要修改 MySQL 配置：
	// echo "[mysqld]
	// bind-address = 0.0.0.0" >> /etc/mysql/conf.d/docker.cnf
	// 重启mysql
	// docker restart mysql_server

	// notejs
}
