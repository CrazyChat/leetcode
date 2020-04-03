package structure

import "errors"

type KeyType = int
type ValType = int

type LruNode struct {
	Key KeyType
	Val ValType
	Pre *LruNode
	Next *LruNode
}

type LruCache struct {
	limit int
	HashMap map[KeyType]*LruNode
	head *LruNode
	end *LruNode
}

func CreateLruCache(cap int) LruCache {
	return LruCache{limit: cap, HashMap:make(map[KeyType]*LruNode, cap)}
}

func (l *LruCache) Get(key KeyType) (ValType, error) {
	if node, exist := l.HashMap[key]; exist {
		l.refreshNode(node)
		return node.Val, nil
	}
	var i ValType
	return i, errors.New("No exist")
}

// Put 存在时修改，不存在时新增
func (l *LruCache) Put(key KeyType, val ValType) {
	node, exists := l.HashMap[key]
	if exists {
		node.Val = val
		l.refreshNode(node)
		return
	}
	// 不存在，插入
	if len(l.HashMap) >= l.limit {
		// todo 删除链尾
		oldKey := l.removeNode(l.end)
		delete(l.HashMap, oldKey)
	}
	new_node := LruNode{Key: key, Val: val}
	l.addNode(&new_node)
	l.HashMap[key] = &new_node
	return
}

func (l *LruCache) refreshNode(node *LruNode) {
	if node == l.head {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LruCache) removeNode(node *LruNode) ValType {
	if node == l.end {
		l.end = l.end.Pre
		l.end.Next = nil
	} else if node == l.head {
		l.head = l.head.Next
		l.head.Pre = nil
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
	return node.Val
}

func (l *LruCache) addNode(node *LruNode) {
	if l.head == nil {
		l.head = node
		l.end = node
		return
	}
	node.Next = l.head
	l.head.Pre = node
	l.head = node
	return
}