package main

/*
146. LRU缓存机制
难度:中等
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
进阶:
你是否可以在 O(1) 时间复杂度内完成这两种操作？
示例:
LRUCache cache = new LRUCache( 2 ); // 2为缓存容量
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得关键字 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得关键字 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// LinkNode ...
type LinkNode struct {
	key   int       // 当cache元素满之后，删除map中当数据可以直接使用key
	value int       // value
	pre   *LinkNode // 双向链表，前置
	next  *LinkNode // 双向链表，后续
}

// LRUCache ..
type LRUCache struct {
	cache map[int]*LinkNode // 数据缓存
	cap   int               // 容量
	head  *LinkNode         // 首节点
	tail  *LinkNode         // 尾节点
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	// go会自动初始化,无须0,0,nil,nil来初始化
	tail := &LinkNode{}
	head := &LinkNode{}

	// 双向链表
	head.next = tail
	tail.pre = head

	return LRUCache{
		cache: make(map[int]*LinkNode),
		cap:   capacity,
		head:  head,
		tail:  tail,
	}
}

// Get ...
func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		// 优先级更新,将该元素移动到链表head位上去
		// nil<-head->next,
		// pre<-node->next
		// 将node到前后关系维护一下,说白了直接删除node
		node.pre.next = node.next
		node.next.pre = node.pre

		tmp := l.head
		// node本身要提到前面，所以他到前置为nil
		node.pre = nil
		// 当前head的pre为node
		tmp.pre = node
		// node的next为当前的head
		node.next = tmp
		l.head = node
		return node.value
	}
	return -1
}

// Put ...
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		// 存在直接移动到head
		// 先将node删除
		node.pre.next = node.next
		node.next.pre = node.pre

		tmp := l.head
		// node本身要提到前面，所以他到前置为nil
		node.pre = nil
		// 当前head的pre为node
		tmp.pre = node
		// node的next为当前的head
		node.next = tmp
		l.head = node
		l.cache[key] = node
	} else {
		// 如果已经超了，那么删除链表tail
		if len(l.cache) == l.cap {
			delete(l.cache, l.tail.key)
			pre := l.tail.pre
			pre.next = nil
			l.tail = pre
		}

		// 不存在直接插入到head
		node := &LinkNode{
			key:   key,
			value: value,
		}
		node.pre = nil
		node.next = l.head
		l.head.pre = node
		l.head = node
		// 缓存node
		l.cache[key] = node
	}
	return
}
