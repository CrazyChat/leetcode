package queue

/*
641. 设计循环双端队列
*/

type MyCircularDeque struct {
	head int
	tail int
	cap  int
	val  []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor2(k int) MyCircularDeque {
	return MyCircularDeque{cap: k + 1, val: make([]int, k+1)}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if (this.tail+1)%this.cap == this.head {
		return false
	}
	this.head = (this.head + this.cap - 1) % this.cap
	this.val[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if (this.tail+1)%this.cap == this.head {
		return false
	}
	this.val[this.tail] = value
	this.tail = (this.tail + 1) % this.cap
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == this.tail {
		return false
	}
	this.head = (this.head + 1) % this.cap
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.head == this.tail {
		return false
	}
	this.tail = (this.tail + this.cap - 1) % this.cap
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.head == this.tail {
		return -1
	}
	return this.val[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.head == this.tail {
		return -1
	}
	return this.val[(this.tail+this.cap-1)%this.cap]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.cap == this.head
}
