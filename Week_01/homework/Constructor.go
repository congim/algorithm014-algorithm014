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
	key   int
	value int
	prev  *LinkNode
	next  *LinkNode
}

// LRUCache ...
type LRUCache struct {
	cap   int
	cache map[int]*LinkNode
	head  *LinkNode
	tail  *LinkNode
}

func newNode(key, value int) *LinkNode {
	return &LinkNode{
		key:   key,
		value: value,
	}
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	lruc := LRUCache{
		cache: map[int]*LinkNode{},
		head:  newNode(0, 0),
		tail:  newNode(0, 0),
		cap:   capacity,
	}

	lruc.head.next = lruc.tail
	lruc.tail.prev = lruc.head
	return lruc
}

// Get ...
func (l *LRUCache) Get(key int) int {
	// 找不到直接返回-1
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	// 找到的话要移动node到head位置
	// 过程如下:
	// 当前位置删除: node.prev.next -- node -- node.next  (node.prev.next = node.next; node.next.prev = node.prev)
	node.prev.next = node.next
	node.next.prev = node.prev

	// 将node移动到head位置
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node

	return node.value
}

// Put ...
func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := newNode(key, value)
		l.cache[key] = node
		// l.addToHead(node)
		// 将node移动到head位置
		node.prev = l.head
		node.next = l.head.next
		l.head.next.prev = node
		l.head.next = node
		if len(l.cache) > l.cap {
			removed := l.removeTail()
			delete(l.cache, removed.key)
		}
	} else {
		node := l.cache[key]
		node.value = value
		// l.moveToHead(node)
		node.prev.next = node.next
		node.next.prev = node.prev

		node.prev = l.head
		node.next = l.head.next
		l.head.next.prev = node
		l.head.next = node

	}
}

func (l *LRUCache) addToHead(node *LinkNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *LinkNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *LinkNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

///////////////////////

// LinkNode ...
// type LinkNode struct {
// 	key   int
// 	value int
// 	prev  *LinkNode
// 	next  *LinkNode
// }

// // LRUCache ...
// type LRUCache struct {
// 	cap   int
// 	cache map[int]*LinkNode
// 	head  *LinkNode
// 	tail  *LinkNode
// }

// func newNode(key, value int) *LinkNode {
// 	return &LinkNode{
// 		key:   key,
// 		value: value,
// 	}
// }

// // Constructor ...
// func Constructor(capacity int) LRUCache {
// 	lruc := LRUCache{
// 		cache: map[int]*LinkNode{},
// 		head:  newNode(0, 0),
// 		tail:  newNode(0, 0),
// 		cap:   capacity,
// 	}

// 	lruc.head.next = lruc.tail
// 	lruc.tail.prev = lruc.head
// 	return lruc
// }

// // Get ...
// func (l *LRUCache) Get(key int) int {
// 	// 找不到直接返回-1
// 	node, ok := l.cache[key]
// 	if !ok {
// 		return -1
// 	}
// 	// 找到的话要移动node到head位置
// 	// 过程如下:
// 	// 当前位置删除: node.prev.next -- node -- node.next  (node.prev.next = node.next; node.next.prev = node.prev)
// 	node.prev.next = node.next
// 	node.next.prev = node.prev

// 	// 将node移动到head位置
// 	// node.prev = l.head
// 	// node.next = l.head.next
// 	// l.head.next.prev = node
// 	// l.head.next = node
// 	l.addToHead(node)
// 	return node.value
// }

// // Put ...
// func (l *LRUCache) Put(key int, value int) {
// 	if _, ok := l.cache[key]; !ok {
// 		node := newNode(key, value)
// 		l.cache[key] = node

// 		node.prev = l.head
// 		node.next = l.head.next
// 		l.head.next.prev = node
// 		l.head.next = node
// 		if len(l.cache) > l.cap {
// 			// 先删除数据
// 			delete(l.cache, l.tail.prev.key)
// 			// 删除节点
// 			l.tail.prev.prev.next = l.tail.prev.next
// 			l.tail.prev = l.tail.prev.prev
// 		}
// 	} else {
// 		node := l.cache[key]
// 		node.value = value
// 		// 移动到头
// 		// node.next = l.head.next
// 		// node.prev = l.head
// 		// l.head.next.prev = node
// 		// l.head.next = node
// 		l.addToHead(node)
// 	}
// }

// func (l *LRUCache) addToHead(node *LinkNode) {
// 	node.prev = l.head
// 	node.next = l.head.next
// 	l.head.next.prev = node
// 	l.head.next = node
// }

// func (l *LRUCache) removeNode(node *LinkNode) {
// 	node.prev.next = node.next
// 	node.next.prev = node.prev
// }

// func (l *LRUCache) moveToHead(node *LinkNode) {
// 	l.removeNode(node)
// 	l.addToHead(node)
// }

// func (l *LRUCache) removeTail() *LinkNode {
// 	node := l.tail.prev
// 	l.removeNode(node)
// 	return node
// }
///////////////////////
// // LinkNode ...
// type LinkNode struct {
// 	key   int       // 当cache元素满之后，删除map中当数据可以直接使用key
// 	value int       // value
// 	pre   *LinkNode // 双向链表，前置
// 	next  *LinkNode // 双向链表，后续
// }

// // LRUCache ..
// type LRUCache struct {
// 	cache map[int]*LinkNode // 数据缓存
// 	cap   int               // 容量
// 	head  *LinkNode         // 首节点
// 	tail  *LinkNode         // 尾节点
// }

// // Constructor ...
// func Constructor(capacity int) LRUCache {
// 	// go会自动初始化,无须0,0,nil,nil来初始化
// 	tail := &LinkNode{}
// 	head := &LinkNode{}

// 	// 双向链表
// 	head.next = tail
// 	tail.pre = head

// 	return LRUCache{
// 		cache: make(map[int]*LinkNode),
// 		cap:   capacity,
// 		head:  head,
// 		tail:  tail,
// 	}
// }

// // Get ...
// func (l *LRUCache) Get(key int) int {
// 	if node, ok := l.cache[key]; ok {
// 		// 优先级更新,将该元素移动到链表head位上去
// 		// nil<-head->next,
// 		// pre<-node->next
// 		// 将node到前后关系维护一下,说白了直接删除node
// 		node.pre.next = node.next
// 		node.next.pre = node.pre

// 		tmp := l.head
// 		// node本身要提到前面，所以他到前置为nil
// 		node.pre = nil
// 		// 当前head的pre为node
// 		tmp.pre = node
// 		// node的next为当前的head
// 		node.next = tmp
// 		l.head = node
// 		return node.value
// 	}
// 	return -1
// }

// // Put ...
// func (l *LRUCache) Put(key int, value int) {
// 	if node, ok := l.cache[key]; ok {
// 		node.value = value
// 		// 存在直接移动到head
// 		// 先将node删除
// 		node.pre.next = node.next
// 		node.next.pre = node.pre

// 		tmp := l.head
// 		// node本身要提到前面，所以他到前置为nil
// 		node.pre = nil
// 		// 当前head的pre为node
// 		tmp.pre = node
// 		// node的next为当前的head
// 		node.next = tmp
// 		l.head = node
// 		l.cache[key] = node
// 	} else {
// 		// 如果已经超了，那么删除链表tail
// 		if len(l.cache) > l.cap {
// 			delete(l.cache, l.tail.key)
// 			pre := l.tail.pre
// 			pre.next = nil
// 			l.tail = pre
// 		}

// 		// 不存在直接插入到head
// 		node := &LinkNode{
// 			key:   key,
// 			value: value,
// 		}
// 		node.pre = nil
// 		node.next = l.head
// 		l.head.pre = node
// 		l.head = node
// 		// 缓存node
// 		l.cache[key] = node
// 	}
// 	return
// }
