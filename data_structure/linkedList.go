package data_structure

type llnode[T any] struct {
	next *llnode[T]
	prev *llnode[T]
	val  T
}

type LinkedList[T any] struct {
	head *llnode[T]
	tail *llnode[T]
	len  int
}

func (l *LinkedList[T]) PushTop(val T) {
	var cur *llnode[T]
	cur.val = val
	if l.len == 0 {
		l.head = cur
		l.tail = cur
	} else {
		l.head.prev = cur
		cur.next = l.head
		l.head = cur
	}
	l.len += 1
}

func (l *LinkedList[T]) PushTail(val T) {
	var cur *llnode[T]
	cur.val = val
	if l.len == 0 {
		l.head = cur
		l.tail = cur
	} else {
		l.tail.next = cur
		cur.prev = l.tail
		l.tail = cur
	}
	l.len += 1
}

func (l *LinkedList[T]) PopTop() T {
	if l.len == 0 {
		panic(any("end of stack"))
	}
	tmp := l.head.val
	l.head = l.head.next
	l.head.prev = nil
	return tmp
}

func (l *LinkedList[T]) PopTail() T {
	if l.len == 0 {
		panic(any("end of stack"))
	}
	tmp := l.tail.val
	l.tail = l.tail.prev
	l.tail.next = nil
	return tmp
}
