package cache

import (
	"container/list"
)

/*

problem : https://leetcode.com/problems/lru-cache/

use a hashmap to store key/value
use a doubly linked-list (DLL) to store the order each key is accessed.
when Put:
	if HashMap[key] already exists => HashMap[key] = value
	else
		- check size of Hashmap
			if size == capacity
				=>  pop the key at the head of DLL and delete it in Hashmap

		- HashMap[key] = value; push back this key to DDL

when Get:
	- if HashMap[key] is not exist => return -1
	- else
		remove it in DLL, then add it to back of DLL

*/

type LRUCache struct {
	capacity int
	ddl      *list.List
	itemMap  map[int]*list.Element
}

type KeyValue struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ddl:      list.New(),
		itemMap:  make(map[int]*list.Element),
	}
}

func (this *LRUCache) Put(key int, value int) {

	//if HashMap[key] already exists => HashMap[key] = value
	if _, exist := this.itemMap[key]; exist {

		item := this.itemMap[key]

		//update value
		item.Value = KeyValue{
			key:   key,
			value: value,
		}

		//move the item to back of the DDL
		this.ddl.MoveToBack(item)

	} else {
		//	check size of Hashmap:
		//	if size == capacity
		//		=>  pop the key at the head of DLL and delete it in Hashmap
		if len(this.itemMap) == this.capacity {
			item := this.ddl.Front()
			this.ddl.Remove(item)
			delete(this.itemMap, item.Value.(KeyValue).key)
		}

		//- HashMap[key] = value; push back this key to DDL
		this.ddl.PushBack(KeyValue{
			key:   key,
			value: value,
		})
		item := this.ddl.Back()
		this.itemMap[key] = item
	}
}

func (this *LRUCache) Get(key int) int {
	item, exist := this.itemMap[key]
	if !exist {
		return -1
	}

	this.ddl.MoveToBack(item)
	res := item.Value.(KeyValue)
	return res.value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//2 3 xxx
//1 1
//4 1
