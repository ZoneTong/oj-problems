package main

import "fmt"

func main() {
	q := NewDeque()
	n1 := &Node{value: 1}
	q.Insert(nil, n1)
	fmt.Println(q)

	n2 := &Node{value: 2}
	q.Insert(n1, n2)
	fmt.Println(q)

	n3 := &Node{value: 3}
	q.Insert(n1, n3)
	fmt.Println(q)

	q.Del(n3)
	fmt.Println(q)
	q.Del(n1)
	fmt.Println(q)
}
