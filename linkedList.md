##### 反转链表
```go
func reverse(head *listNode) *listNode {
	var pre *listNode // 头结点的 next
	for head != nil {
		tmp := head.next 
		head.next = pre // next 指向上一个节点
		pre = head // pre = 当前节点
		head = tmp // head = next 下一轮循序
	}
	return pre
}
```
##### 链表区间反转
```go
func reverseBetween(node *listNode, m int, n int) *listNode  {
	if m == n || m > n {
		return node
	}
	//node =  1->2->3->4->5->6->7->8 m = 2, n = 7
	// 输出 1->7->6->5->4->3->2->8
	// 首先找到在 链表中找到m 将 m.next 放到 per 和 pre.next 中间 然后 m.next = m.next.next 下一轮循环
	//第一步 找到3 放到 1 和 2中 由  1->2->3->4->5->6->7->8 变到 1->3->2->4->5->6->7->8
	//第二步 找到4 放到 1 和 3中 由 1->3->2->4->5->6->7->8 变到 1->4->3->2->5->6->7->8
	//第三步 找到5 放到 1 和 4中 由 1->4->3->2->5->6->7->8 变到 1->5->4->3->2->6->7->8

	// 先找到 2之前的位置
	pre := node
	for i := 1; i < m - 1; i++ {
		pre = pre.next
	}
	cur := pre.next // 节点 2 的位置

	for i := m; i < n; i++ {
		tmp := cur.next // 第一轮 节点 3
		cur.next = tmp.next // 第一轮 节点 2 指向 4
		tmp.next = pre.next // 第一轮 节点 3 指向 2
		pre.next = tmp// 第一轮 节点 1 指向 3
	}
	return node
}
```

##### 检测链表是否有环
```go

func hasRepeatByMap(head *listNode) bool {

	seen := map[*listNode]struct{}{}
	for head != nil {
		// 如果map存在则有环
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.next
	}
	return false
}


func hasRepeat(head *listNode) bool {

	// 首先定义两个指针 slow 和 fast，初始都指向链表头指针 head；
	// 然后，slow 每走一步（每次移动一个节点），fast 走两步（移动两个节点）；
	// 如果 fast 遇到了 NULL，则链表不存在环；
	// 如果 slow==fast，则链表存在环；
	
	/*
		如果存在环，fast 和 slow 肯定会先后进入环；
		当 slow 刚好进入环时，fast 肯定已经在环内的某一个位置了；
		而又因 fast 走的速度是 slow 的两倍，所以在 slow 在环内遍历一遍的过程中，fast 指针必然会与 slow 相遇。
		其实两个指针就像是两个人在田径场上跑步，一个人的速度是另一个人速度的两倍。无论两个人的起点在何处，慢的人跑完一圈，快的人必然跑完两圈，所以他们必然相遇过。
	*/

	slow := head
	fast:=  head

	for fast != nil && fast.next != nil {
		fast = head.next.next
		slow = head.next
		if slow == fast {
			return true
		}
	}
	return false
}
```
