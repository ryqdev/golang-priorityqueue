package golang_priorityqueue

type PriorityQueue[T comparable] struct {
	Array []T
	Rule  func(T, T) bool
}

func (p *PriorityQueue[T]) Heapify() {
	for i := len(p.Array)/2 - 1; i >= 0; i-- {
		p.heapifyDown(i)
	}
}

func (p *PriorityQueue[T]) heapifyDown(idx int) {
	cur := idx
	left := 2*idx + 1
	right := 2*idx + 2
	if left < len(p.Array) && p.Rule(p.Array[left], p.Array[cur]) {
		cur = left
	}
	if right < len(p.Array) && p.Rule(p.Array[right], p.Array[cur]) {
		cur = right
	}
	if cur != idx {
		p.Array[cur], p.Array[idx] = p.Array[idx], p.Array[cur]
		p.heapifyDown(cur)
	}
}

func (p *PriorityQueue[T]) Push(val T) {
	p.Array = append(p.Array, val)
	p.heapifyUp(len(p.Array) - 1)
}

func (p *PriorityQueue[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	cur := idx
	father := (idx - 1) / 2
	if p.Rule(p.Array[cur], p.Array[father]) {
		cur = father
	}
	if cur != idx {
		p.Array[cur], p.Array[idx] = p.Array[idx], p.Array[cur]
		p.heapifyUp(cur)
	}
}

func (p *PriorityQueue[T]) Top() T {
	if p.Empty() {
		panic("priority queue is empty")
	}
	return p.Array[0]
}

func (p *PriorityQueue[T]) Pop() {
	if p.Empty() {
		panic("priority queue is empty")
	}
	p.Array[0], p.Array[len(p.Array)-1] = p.Array[len(p.Array)-1], p.Array[0]
	p.Array = p.Array[:len(p.Array)-1]
	p.heapifyDown(0)
}

func (p *PriorityQueue[T]) Empty() bool {
	return len(p.Array) == 0
}
