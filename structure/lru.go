package structure

/*
双向队列保存 LRUNode 的指针
从头到尾使用频率：从高到低
即队尾为最久未使用
*/
type LRUCache struct {
	cap int
	head *LRUNode
	end *LRUNode
	dataMap map[int]*LRUNode
}

type LRUNode struct {
	key int
	val int
	pre *LRUNode
	next *LRUNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity, &LRUNode{}, &LRUNode{}, make(map[int]*LRUNode, capacity)}
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
	newNode := LRUNode{key: key, val: value}
	if len(this.dataMap) >= this.cap {
		delete(this.dataMap, this.end.pre.key)
		this.remove(this.end.pre)
	}
	this.putHead(&newNode)
	this.dataMap[key] = &newNode
}

func (this *LRUCache) putHead(node *LRUNode) {
	// 新增节点和head.next关联起来
	this.head.next.pre = node
	node.next = this.head.next
	// head和新增加点关联起来
	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) remove(node *LRUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}