package main

func main() {
	// mysql 的 悲观锁
	// 行锁 for update
	// select * from table1 where age=18 for update
	// 如果 where 查询字段 age 有索引，是行锁，只会锁住当前行，其他 for update 无法查询修改
	// 如果 age 没有索引，行锁会升级成表锁，整张表无法查询修改
	// 如果 锁 的查询没有查询到结果，不会锁，但是没有索引的情况下仍然会锁住表

	// mysql 乐观锁
	// 其实是不加锁， 而已用 where version=version 来判断更新的数据是不是更新前获取的那一条，更新完成 version 加 1

	// 共享锁（读锁，多个事务可同时读） LOCK IN SHARE MODE;
	// START TRANSACTION;
	// SELECT * FROM users WHERE id = 1 LOCK IN SHARE MODE;
	// -- 其他事务仍可以读取 id = 1，但不能修改或删除
	// COMMIT;

	// 排他锁（写锁）FOR UPDATE;
	// START TRANSACTION;
	// SELECT * FROM users WHERE id = 1 FOR UPDATE;
	// -- 其他事务不能读取、修改或删除 id = 1
	// COMMIT;
	//
	//
}
