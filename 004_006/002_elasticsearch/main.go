package main

func main() {
	// 1. 索引（Index）
	//	•	MySQL：
	// 在 MySQL 中，数据通常存储在数据库（Database）下的一个或多个表（Table）中。
	//	•	Elasticsearch：
	// Elasticsearch 的 Index 类似于一个数据库或表，但更偏向于“表”的概念。每个索引存储一组文档（Document），而索引内部的文档具有相同的结构（通过 Mapping 定义）。
	//	•	对比说明：
	//	•	可以将一个 Elasticsearch 索引看作 MySQL 中的一个表。
	//	•	索引中的 Mapping 相当于 MySQL 的表结构（列定义）。
	//
	// ⸻
	//
	// 2. 文档（Document）
	//	•	MySQL：
	// 数据以行（Row）的形式存储在表中，每行代表一条记录。
	//	•	Elasticsearch：
	// Document 是 Elasticsearch 中存储数据的基本单元，通常以 JSON 格式存储。
	//	•	对比说明：
	//	•	一个 Document 类似于 MySQL 表中的一行。
	//	•	每个文档由多个字段（Field）组成，对应于表中的列（Column）。
	//
	// ⸻
	//
	// 3. 字段（Field）
	//	•	MySQL：
	// 表的每一列定义了数据类型和约束。
	//	•	Elasticsearch：
	// Field 是文档中的属性，通过 Mapping 定义其数据类型、格式及如何被索引和分析。
	//	•	对比说明：
	//	•	字段与 MySQL 的列类似，但 Elasticsearch 的字段不仅决定存储格式，还决定了如何进行全文搜索和分词等操作。
	//
	// ⸻
	//
	// 4. Mapping（映射）
	//	•	MySQL：
	// 建表语句定义了表的结构和各列的数据类型。
	//	•	Elasticsearch：
	// Mapping 定义了索引中各个字段的数据类型、格式和分词规则等信息。
	//	•	对比说明：
	//	•	Mapping 类似于 MySQL 的表结构定义，不过 Elasticsearch 更灵活，支持动态映射（可以在索引时自动推导字段类型），也支持手动定义以获得更精准的搜索行为。
	//
	// ⸻
	//
	// 5. 分片（Shard）与副本（Replica）
	//	•	MySQL：
	// MySQL 通常运行在单一服务器或主从复制、分布式数据库（如分库分表）方案下，但本身不直接管理数据分片。
	//	•	Elasticsearch：
	//	•	Shard： 一个索引可以被分为多个分片，每个分片是一个 Lucene 索引，可以分布在不同节点上，从而实现水平扩展。
	//	•	Replica： 分片的副本，用于提高数据的高可用性和搜索吞吐量。
	//	•	对比说明：
	//	•	你可以把分片看作是 MySQL 表的物理拆分，副本则类似于 MySQL 的主从复制或读写分离的备份机制，但它们是 Elasticsearch 内置、透明管理的。
	//
	// ⸻
	//
	// 6. 查询语言与索引机制
	//	•	MySQL：
	// 使用 SQL 语言，依靠 B-Tree 索引等传统索引结构来快速查找数据。
	//	•	Elasticsearch：
	// 使用基于 JSON 的 Query DSL 进行查询，底层构建了倒排索引（Inverted Index）来支持高效的全文搜索。
	//	•	对比说明：
	//	•	虽然两者都支持索引，但 Elasticsearch 的倒排索引专门优化了对文本内容的快速搜索和相关性排序，而 MySQL 的索引更侧重于精确匹配和范围查询。
	//
	// ⸻
	//
	// 7. 分布式特性
	//	•	MySQL：
	// 单机部署时一般不具备原生分布式能力，分布式 MySQL 通常需要借助额外的中间件或方案（如 MySQL Cluster、分库分表方案）。
	//	•	Elasticsearch：
	// 天然设计为分布式系统，支持数据分片、副本、自动故障转移和集群管理，适用于处理海量数据的实时搜索和分析。
	//	•	对比说明：
	//	•	Elasticsearch 在分布式扩展、横向扩展和高可用性方面要比传统 MySQL 更容易实现，并且管理上更为透明和自动化。
	//
	// ⸻
	//
	// 总结
	//	•	索引/表： Elasticsearch 索引类似于 MySQL 的表；Mapping 类似于表结构定义。
	//	•	文档/记录： Document 就像表中的一行记录，每个字段对应一个列。
	//	•	分片与副本： Elasticsearch 内置分片与副本机制，类似于 MySQL 的分库分表和复制，但更适合大规模分布式场景。
	//	•	查询机制： Elasticsearch 倒排索引和 Query DSL 主要用于全文搜索和分析，而 MySQL 的 SQL 更适合结构化查询和事务处理。
	//	•	分布式特性： Elasticsearch 从设计上就是一个分布式系统，易于水平扩展，而 MySQL 则需要额外方案来实现分布式部署。
	//
	// 通过这样的类比，如果你熟悉 MySQL 的数据库、表、行、列以及索引、复制等概念，就可以更容易理解 Elasticsearch 的设计理念和基本架构。两者虽各有侧重，但在数据组织和查询索引方面有相似之处，同时 Elasticsearch 专注于高效的全文搜索和分布式扩展能力。

	//

}
