package main

import "fmt"

func main() {
	customerA := &CustomerA{}
	customerB := &CustomerB{}
	
	office := NewsOffice{}
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	office.newspaperCome()
	
	firstUniqChar("leetcode")
}

func firstUniqChar(s string) byte {
	sbyte := []byte(s)
	sMap := make(map[byte]int, len(sbyte))
	for _, v := range sbyte {
		if _, ok := sMap[v]; !ok {
			sMap[v] = 0
		} else {
			sMap[v] += 1
		}
		fmt.Println(sMap)
	}
	for _, v := range sbyte {
		if sMap[v] < 1 {
			return v
		}
	}
	return ' '
}

// 报社 - 客户
type Customer interface {
	update()
}

type CustomerA struct {}

type CustomerB struct {}

func (*CustomerA) update() {
	fmt.Println("我是客户A，我收到报纸了")
}

func (*CustomerB) update() {
	fmt.Println("我是客户B，我收到报纸了")
}

// 报社（被观察者）
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newspaperCome() {
	// 通知所有客户新闻
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}

