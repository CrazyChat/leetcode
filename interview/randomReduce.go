package interview

/*
360 - 随机减法
 */

//func main() {
//	n, m := 0, 0
//	_, _ = fmt.Scanf("%d %d", &n, &m)
//	// 存储选手power值，循环队列
//	power := NewQueue(n+1)
//	for i := 0; i < n; i++ {
//		var temp int
//		_, _ = fmt.Scanf("%d", &temp)
//		power.Enqueue(temp)
//	}
//	// 特例判断
//	if n <= 1 {
//		fmt.Println(0)
//		return
//	}
//	if n == 2 {
//		fmt.Println(1)
//		return
//	}
//	// 擂台
//	play := []int{0, 0}
//	play[0] = power.Dequeue()
//	winTimes := 0
//	result := 0
//	// 循环比赛
//	for winTimes < m {
//		play[1] = power.Dequeue()
//		if isBig(play[0], play[1]) {
//			power.Enqueue(play[1])
//			winTimes++
//		} else {
//			power.Enqueue(play[0])
//			play[0] = play[1]
//			winTimes = 1
//		}
//		result++
//	}
//	fmt.Println(result)
//}
//
//func isBig(x, y int) bool {
//	if x > y {
//		return true
//	}
//	return false
//}
//
//type Queue struct {
//	cap int
//	head int
//	tail int
//	data []int
//}
//// NewQueue 创建一个cap大小的队列
//func NewQueue(capacity int) Queue {
//	return Queue{cap: capacity, data: make([]int, capacity)}
//}
//// Enqueue 入队
//func (q *Queue) Enqueue(value int) {
//	// (Tail + 1) % Cap == Head 表示队列满了(最后一个不能放值)
//	//if (q.tail+1) % q.cap == q.head {
//	//	return errors.New("the queue is full")
//	//}
//	q.data[q.tail] = value
//	q.tail = (q.tail+1) % q.cap
//}
//// Dequeue 出队
//func (q *Queue) Dequeue() int {
//	temp := q.data[q.head]
//	q.head = (q.head + 1) % q.cap
//	return temp
//}
