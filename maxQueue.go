package main
/**
剑指 Offer 59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

限制：
1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]

*/

import "container/list"

func main() {
	obj := Constructor()
	obj.Push_back(1)
	obj.Push_back(6)
	obj.Push_back(7)
	obj.Push_back(7)
	obj.Push_back(5)
	obj.Push_back(4)
}
//----------使用[]int实现----------
type MaxQueue struct {
	q []int
	m []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q: make([]int, 0),
		m: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.m) == 0 {
		return -1
	}
	return this.m[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)

	for len(this.m) > 0 && this.m[len(this.m)-1] < value {
		this.m = this.m[0:len(this.m)-1]
	}
	this.m = append(this.m, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q) > 0 {
		n = this.q[0]
		this.q = this.q[1:]
		if n == this.m[0] {
			this.m = this.m[1:]
		}
	}
	return n
}

//----------使用list实现----------
type MaxQueue_list struct {
	q *list.List
	m *list.List
}

func Constructor_list() MaxQueue_list {
	return MaxQueue_list{
		q: new(list.List),
		m: new(list.List),
	}
}

func (this *MaxQueue_list) Max_value() int {
	if this.m.Len() == 0 {
		return -1
	}
	return this.m.Front().Value.(int)
}

func (this *MaxQueue_list) Push_back(value int) {
	this.q.PushBack(value)
	for this.m.Len() > 0 && this.m.Back().Value.(int) < value {
		this.m.Remove(this.m.Back())
	}
	this.m.PushBack(value)
}

func (this *MaxQueue_list) Pop_front() int {
	n := -1
	if this.q.Len() > 0 {
		n = this.q.Remove(this.q.Front()).(int)
		if n == this.m.Front().Value.(int) {
			this.m.Remove(this.m.Front())
		}
	}
	return n
}