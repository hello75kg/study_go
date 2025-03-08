package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list
	//链表
	//适合插入删除，不适合查询
	var llst list.List

	llst.PushBack("go1")
	llst.PushBack("go2")
	llst.PushBack("go3")
	llst.PushFront("fo1")
	fmt.Println(llst)

	for i := llst.Front(); i != nil; i = i.Next() {
		fmt.Println(i, i.Value)
	}

	for i := llst.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	//插入删除
	var lst list.List
	lst.PushBack("1")
	lst.PushBack("2")
	lst.PushBack("3")
	lst.PushBack("4")
	lst.PushBack("5")

	fmt.Println("----------")
	head := lst.Front()
	lst.Remove(head.Next())
	for i := lst.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("----------")
	lst.InsertBefore("0", head)
	lst.InsertAfter("6", lst.Back())
	for i := lst.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
