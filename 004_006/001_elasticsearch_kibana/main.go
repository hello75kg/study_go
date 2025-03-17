package main

func main() {
	// elasticsearch
	// mysql 搜索问题：性能底/没有相关性排名/无法全文搜索/搜索不准确没有分词
	// 1. Elasticsearch
	//	•	Elasticsearch 是一个基于 Apache Lucene 构建的开源分布式搜索与分析引擎，主要用于全文搜索、日志分析和实时数据分析。
	// 		它提供了一个 RESTful API，能通过简单的 HTTP 请求实现数据的索引、搜索和聚合操作
	// 2. 核心概念
	//	•	集群（Cluster）：由多个节点（Node）组成，协同工作以实现数据的分布式存储与搜索。
	//	•	节点（Node）：集群中的一个单独服务器。每个节点都承担数据存储、搜索、索引等任务。
	//	•	索引（Index）：类似于数据库，存储具有相同特征的数据集合。一个索引由多个文档（Document）构成。
	//	•	文档（Document）：数据的基本单元，通常以 JSON 格式存储。每个文档都属于某个索引。
	//	•	分片（Shard）：为了实现横向扩展，索引会被拆分成多个分片。每个分片可以存储在不同的节点上。
	//	•	副本（Replica）：分片的拷贝，用于提高数据的可用性和搜索性能。
	// 3. 使用场景
	//	•	全文搜索：构建搜索引擎，支持高效的模糊搜索、相关性排序等。
	//	•	日志与指标分析：通过 ELK（Elasticsearch、Logstash、Kibana）或 EFK（Elasticsearch、Fluentd、Kibana）方案，对大规模日志数据进行实时分析与可视化。
	//	•	实时分析：对海量数据进行快速聚合计算和数据分析。
	// 4. 安装方式
	// (1) 官方二进制安装
	//	•	访问 Elasticsearch 官方下载页面 下载最新版本的安装包。
	//	•	解压后，根据平台（Linux、Windows 或 macOS）启动 Elasticsearch。
	// (2) Docker 部署
	//	•	使用 Docker 镜像快速启动：
	// 			docker pull docker.elastic.co/elasticsearch/elasticsearch
	// 			docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch
	// 		以单节点模式启动 Elasticsearch，并映射 9200（HTTP 接口）和 9300（集群通信端口）
	// (3) 集群部署
	//	•	在生产环境下，通常需要部署成集群模式，需要配置节点间的通信、分片、副本等
	// 5. 常用工具与生态
	//	•	Kibana：用于 Elasticsearch 数据的可视化和管理，支持图表、仪表板构建以及实时监控。
	//	•	Logstash / Beats / Fluentd：用于日志数据采集、传输和处理，与 Elasticsearch 协同构建日志分析平台。
	//	•	Elastic Stack：通常指 Elasticsearch、Kibana、Logstash 及 Beats 的组合，形成完整的日志和数据分析解决方案。
	// 6. API 与操作
	//
	// Elasticsearch 提供了一套丰富的 RESTful API，例如：
	//	•	索引文档：PUT /{index}/_doc/{id} 将文档索引到指定索引中。
	//	•	搜索：GET /{index}/_search 执行全文搜索或聚合操作。
	//	•	聚合：在搜索 API 中支持统计、分组、直方图等多种聚合查询。
	//
	// 总结
	// Elasticsearch 作为一个分布式搜索和分析引擎，具有以下特点：
	//	•	高性能搜索：通过倒排索引技术实现高效的全文搜索。
	//	•	分布式架构：支持横向扩展，能够处理海量数据存储和搜索需求。
	//	•	实时数据分析：通过聚合和实时数据处理，适合日志、监控和大数据场景。
	//	•	丰富生态：与 Kibana、Logstash 等工具配合，构建强大的数据分析平台。
	//
	// 这种架构使 Elasticsearch 成为许多企业构建搜索引擎、日志分析平台和实时数据处理系统的重要组件。

	/*

			 mkdir -p ~/.es/data/elasticsearch/data
			 mkdir -p ~/.es/data/elasticsearch/plugins
			 chmod -R 777 ~/.es/data/elasticsearch

			 docker run --name elasticsearch -p 9200:9200 -p 9300:9300
					 -e "discovery.type=single-node" -e ES_JAVA_OPS="-Xms128m -Xmx256m" \
					# -e "xpack.security.enabled=false" \ # 关闭后端密码验证
			         -v /Users/wwj/.es/data/elasticsearch/data:/usr/share/elasticsearch/data \
			         -v /Users/wwj/.es/data/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
			         -d docker.elastic.co/elasticsearch/elasticsearch:8.5.0

			 docker update --restart=always xxx

			 mac的后台登录账号密码： elastic / GaeERczhQ*h_Lh8UmLYa

		// 安装 kibana
		// docker run -d --name kibana -e ELASTICSEARCH_HOSTS="http://192.168.0.249:9200" -p 5601:5601 docker.elastic.co/kibana/kibana:8.5.0

	*/
}
