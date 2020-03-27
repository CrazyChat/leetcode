package structure

import "errors"

type Queue []interface{}

func CreateQueue() *Queue {
	var q = Queue{}
	return &q
}

func (queue *Queue) Push(value interface{}) {
	*queue = append(*queue, value)
}

func (queue Queue) Len() int {
	return len(queue)
}

func (queue Queue) Cap() int {
	return cap(queue)
}

func (queue Queue) Front() (interface{}, error) {
	if len(queue) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return queue[len(queue)-1], nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	theStack := *queue
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[0]
	*queue = theStack[1:]
	return value, nil
}

func (queue Queue) IsEmpty() bool {
	return len(queue) == 0
}