package structure

type LRUCache struct {
	cap int
	head *LruNode
	end *LruNode
	dataMap map[int]*LruNode
}

type LruNode struct {
	key int
	val int
	pre *LruNode
	next *LruNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity, &LruNode{}, &LruNode{}, make(map[int]*LruNode, capacity)}
	cache.head.next = cache.end
	cache.end.pre = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.dataMap[key]; exists {
		this.remove(node)
		this.putHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	/* 存在
	1. 更新
	2. 移动到头部
	*/
	if node, exists := this.dataMap[key]; exists {
		node.val = value
		this.remove(node)
		this.putHead(node)
		return
	}
	/* 不存在
	1. 如果满了，先删除end
	2. 插入到头部和dataMap
	*/
	newNode := LruNode{key: key, val: value}
	if len(this.dataMap) >= this.cap {
		delete(this.dataMap, this.end.pre.key)
		this.remove(this.end.pre)
	}
	this.putHead(&newNode)
	this.dataMap[key] = &newNode
}

func (this *LRUCache) putHead(node *LruNode) {
	// 新增节点和head.next关联起来
	this.head.next.pre = node
	node.next = this.head.next
	// head和新增加点关联起来
	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) remove(node *LruNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}