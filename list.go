package z01dakar

type node struct {
	data interface{}
	next *node
}

type list struct {
	head, tail *node
}

// ListPushBack add a data at the end on a list.
func ListPushBack(l list, data interface{}) list {
	temp := &node{data: data}
	if l.head == nil {
		l.head = temp
		l.tail = temp
	} else if l.head == l.tail {
		l.head.next = temp
		l.tail = temp
	} else {
		l.tail.next = temp
		l.tail = l.tail.next
	}
	return l
}
