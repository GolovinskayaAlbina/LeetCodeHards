package hard

type node struct {
	next *node
	prev *node
	num  int
}

type StackMedianFinder struct {
	top *node
	len int
}

func StackConstructor() MedianFinder {
	return MedianFinder{}
}

func (f *StackMedianFinder) StackAddNum(num int) {
	n := &node{
		num: num,
	}
	if f.len == 0 {
		f.top = n
		//f.median = n
		f.len += 1
		return
	} else {
		f.len += 1
		if num <= f.top.num { //add to top
			f.top.prev = n
			n.next = f.top
			f.top = n
		} else { // try to find the place
			tmp := f.top
			for tmp.next != nil && tmp.next.num <= num {
				tmp = tmp.next
			}
			if tmp.next == nil { // add to the end
				tmp.next = n
				n.prev = tmp
			} else { // add in the middle
				tmp.next.prev = n
				n.next = tmp.next
				tmp.next = n
			}
		}
	}
}

func (f *StackMedianFinder) StackFindMedian() float64 {
	if f.len == 0 {
		return 0
	}
	cur := f.len / 2
	if f.len%2 == 0 {
		cur = cur - 1
	}
	tmp := f.top
	for cur > 0 {
		tmp = tmp.next
		cur--
	}
	if f.len%2 == 0 {
		return float64(tmp.num+tmp.next.num) / 2
	}
	return float64(tmp.num)
}
