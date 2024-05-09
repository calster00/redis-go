package store

type HashMap map[string]string

type LinkedList struct {
	head, tail *element
}

type element struct {
	next *element
	prev *element
	val  string
}

func (l *LinkedList) AddFirst(val string) {
	el := &element{val: val, next: l.head}

	if l.head == nil {
		l.head = el
		l.tail = l.head
	} else {
		l.head.prev = el
		l.head = el
	}
}

func (l *LinkedList) AddLast(val string) {
	el := &element{val: val, prev: l.tail}

	if l.tail == nil {
		l.head = el
		l.tail = l.head
	} else {
		l.tail.next = el
		l.tail = el
	}
}

func (l *LinkedList) RemoveFirst() (el element) {
	if l.head == nil {
		return el
	}
	el = *l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	return el
}

func (l *LinkedList) RemoveLast() (el element) {
	if l.tail == nil {
		return el
	}
	el = *l.tail
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	return el
}
