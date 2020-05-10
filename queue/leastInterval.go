package queue

import "sort"

/*
621. 任务调度器
法 1. 排序后执行法
法 2. 优先队列
法 3. 设计
*/

func LeastInterval(tasks []byte, n int) int {
	tasksTimes := make([]int, 26)
	for _, v := range tasks {
		tasksTimes[v-'A']++
	}
	sort.Ints(tasksTimes)
	maxTimes := tasksTimes[25] - 1
	vacancyNums := maxTimes * n
	for i := 24; i >= 0; i-- {
		if tasksTimes[i] < maxTimes {
			vacancyNums -= tasksTimes[i]
		} else {
			vacancyNums -= maxTimes
		}
	}
	if vacancyNums > 0 {
		return vacancyNums + len(tasks)
	}
	return len(tasks)
}

//func LeastInterval(tasks []byte, n int) int {
//	tasksTimes := make([]int, 26)
//	// 统计任务出现次数
//	for _, v := range tasks {
//		tasksTimes[v-'A']++
//	}
//	sort.Ints(tasksTimes)
//	fmt.Println(tasksTimes)
//	// 排序完成后最多任务的在下标25处，任务是A还是B还是C等已经不重要了
//	resultTimes := 0
//	// n+1 为一轮
//	for tasksTimes[25] > 0 {
//		i := 0
//		// 执行n+1次(任务+待命)
//		for i <= n {
//			if tasksTimes[25] == 0 {
//				break
//			}
//			if i < 26 && tasksTimes[25-i] > 0 {
//				tasksTimes[25-i]--
//			}
//			i++
//			resultTimes++
//		}
//		sort.Ints(tasksTimes)
//		fmt.Println(tasksTimes)
//	}
//	return resultTimes
//}
