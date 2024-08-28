package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 声明一个空列表
	l := list.New()

	// 尾部插入
	l.PushBack("canon") // canon
	// 头部插入
	l.PushFront(67) // 67,canon

	// 尾部添加后保存元素句柄
	element := l.PushBack("first") // 67,canon,first

	// 在first后面添加high
	l.InsertAfter("high", element) // 67,canon,first,high

	// 在first之前添加noon
	l.InsertBefore("noon", element) // 67,canon,noon,first,high

	// 使用
	l.Remove(element) // 67,canon,noon,high

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
