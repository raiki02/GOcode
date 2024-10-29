// // package main

// // import "fmt"

// // type Node struct {
// // 	val  int
// // 	next *Node
// // }

// // type List struct {
// // 	Node
// // 	head *Node
// // }

// // func (l List) show() {
// // 	if l.head.next == nil {
// // 		fmt.Println("list is empty")
// // 		return
// // 	}

// // 	p := l.head
// // 	for p != nil {
// // 		fmt.Println(*p, " ")
// // 		p = p.next
// // 	}
// // 	fmt.Print("\n")
// // }

// // func (l List) push_back(val int) {
// // 	// p := l.head
// // 	// for p.next != nil {
// // 	// 	p = p.next
// // 	// }
// // 	// p.next = new(Node)
// // 	// p.next.val = val

// // 	if l.head.next == nil {
// // 		l.head.next.val = val
// // 	}
// // 	p := l.head
// // 	for p.next != nil {
// // 		p = p.next
// // 	}
// // 	newNode := Node{val, nil}
// // 	p.next = &newNode
// // 	//*(p.next).val = val

// // }

// // func main() {
// // 	var l List
// // 	l.push_back(1)
// // 	l.push_back(12)
// // 	l.push_back(25)
// // 	l.push_back(2)
// // 	l.push_back(36)
// // 	l.push_back(79)
// // 	l.show()
// // }

// package main

// import "fmt"

// type Node struct {
// 	val  int
// 	next *Node
// }

// type List struct {
// 	Node
// 	head *Node
// }

// func (l List) show() {
// 	p := l.head.next
// 	for p != nil {
// 		fmt.Print(p.val, " ")
// 		p = p.next
// 	}
// 	fmt.Print("\n")
// }

// func (l List) push_back(val int) {
// 	newNode := Node{val, nil}
// 	p := l.head
// 	for p.next != nil {
// 		p = p.next
// 	}
// 	p.next = &newNode

// }

// func (l List) push_front(val int) {
// 	newNode := &Node{val, nil}
// 	p := l.head.next
// 	newNode.next = p
// 	l.head.next = newNode
// }

// func main() {
// 	var l List = List{Node{0, nil}, &Node{}}
// 	//l.head = &Node{}  //这个！！
// 	l.push_back(12)
// 	l.push_back(1)
// 	l.push_back(24)
// 	l.push_back(43)
// 	l.push_back(96)
// 	l.show()

// 	l.push_front(111)
// 	l.push_front(333)
// 	l.push_front(222)
// 	l.show()
// }

//=====================================
// package main

// import "fmt"

// type Node struct {
// 	val  int
// 	next *Node
// }

// type List struct {
// 	Node
// 	head *Node
// }

// func (l List) push_back(val int) {
// 	newNode := &Node{val, nil}
// 	p := l.head
// 	for p.next != nil {
// 		p = p.next
// 	}
// 	p.next = newNode
// }

// func (l List) show() {
// 	p := l.head.next
// 	for p != nil {
// 		fmt.Print(p.val, " ")
// 		p = p.next
// 	}
// 	fmt.Print("\n")
// }

// func main() {
// 	l := List{Node{0, nil}, new(Node)}
// 	l.push_back(1)
// 	l.push_back(1)
// 	l.push_back(4)
// 	l.push_back(5)
// 	l.push_back(1)
// 	l.push_back(4)
// 	l.show()
// }
