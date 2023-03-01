package main

import (
	"testing"
)

/**

 请你设计并实现一个满足
 LRU (最近最少使用) 缓存 约束的数据结构。



 实现
 LRUCache 类：





 LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-
value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。




 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



 示例：


输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4




 提示：


 1 <= capacity <= 3000
 0 <= key <= 10000
 0 <= value <= 10⁵
 最多调用 2 * 10⁵ 次 get 和 put


 Related Topics 设计 哈希表 链表 双向链表 👍 2565 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

type DLinkedNode struct {
	Key  int
	Val  int
	Prev *DLinkedNode
	Next *DLinkedNode
}

type LRUCache struct {
	head     *DLinkedNode
	tail     *DLinkedNode
	cacheMap map[int]*DLinkedNode
	capacity int
	size     int
}

func Constructor(capacity int) LRUCache {

	head := &DLinkedNode{Next: nil}
	tail := &DLinkedNode{Next: nil}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		cacheMap: make(map[int]*DLinkedNode, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// 获取对应的val
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}

	// 如果已经在链表头，就直接忽略
	if this.head.Next == node {
		return node.Val
	}

	// 将node提到链表头
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Prev = nil
	node.Next = nil
}

func (this *LRUCache) addToHead(node *DLinkedNode) {

	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeTail() *DLinkedNode {
	// 删除队尾
	if this.tail.Prev == this.head {
		return nil
	}

	// 删除map
	node := this.tail.Prev
	this.tail.Prev = node.Prev
	node.Prev.Next = this.tail

	return node

}

func (this *LRUCache) Put(key int, value int) {
	// 如果原来有，就修改val, 再放回表头
	node, ok := this.cacheMap[key]
	if ok {
		node.Val = value
		this.moveToHead(node)
		return
	}

	// 检查容量，如果满了，就删掉链表尾部
	if this.capacity == this.size {
		deletedNode := this.removeTail()
		if deletedNode != nil {
			delete(this.cacheMap, deletedNode.Key)
			this.size--
		}
	}

	// 把新节点放到队头
	node = &DLinkedNode{Key: key, Val: value}
	this.addToHead(node)
	this.cacheMap[key] = node
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestLruCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1) // 缓存是 {1=1}
	lruCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lruCache.Get(1)    // 返回 1
	lruCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lruCache.Get(2)    // 返回 -1 (未找到)
	lruCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lruCache.Get(1)    // 返回 -1 (未找到) lRUCache.get(3); // 返回 3
	lruCache.Get(4)    // 返回 4
}
