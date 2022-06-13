package main

import "main/linklist"

func main() {
	sll := linklist.CreateSinglyLinkList[int]()
	sll.Insert(10)
	sll.Insert(20)
	sll.Insert(30)
	sll.DisplayLinkList()
	sll.Delete(10)
	sll.DisplayLinkList()

}
