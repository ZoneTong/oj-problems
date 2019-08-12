package main

import "fmt"

type Node struct {
	value int
	pre   *Node
	next  *Node
}
type Deque Node

func NewDeque() *Deque {

	var head Node
	head.pre = nil
	head.next = nil
	var q = Deque(head)
	return &q
}

/** Descrption: TODO
 *  CreateTime: 2019/08/11 09:52:12
 *      Author: tong.zhou@baishancloud.com
 */
func (q *Deque) Insert(cur, n *Node) {
	// len==0
	// q->head
	// q-> head -> node
	if cur == nil {
		q.next = n
		n.pre = nil
		n.next = nil
	} else {
		next := cur.next
		cur.next = n
		n.pre = cur
		n.next = next
		if next != nil {
			next.pre = n
		}
	}

	// len>0
	// q->head->... node1> node2 >...
	// q->head-> ... node1 > node > node2

}

/** Descrption: TODO
 *  CreateTime: 2019/08/11 10:08:15
 *      Author: tong.zhou@baishancloud.com
 */
func (q *Deque) Del(n *Node) {
	if n == nil {
		return
	}

	if n.pre == nil {
		q.next = n.next
		if n.next != nil {
			n.next.pre = nil
		}
	} else {
		pre := n.pre
		next := n.next
		pre.next = next
		if next != nil {
			next.pre = pre
		}
	}
}

/** Descrption: TODO
 *  CreateTime: 2019/08/11 10:05:03
 *      Author: tong.zhou@baishancloud.com
 */
func (q *Deque) String() string {
	var arr []int
	var n = q.next
	for n != nil {
		arr = append(arr, n.value)
		n = n.next
	}
	return fmt.Sprint(arr)
}
