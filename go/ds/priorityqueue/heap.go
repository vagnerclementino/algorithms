package ds

type Heap struct {
	values        []int
	multiplicator int
}

func NewMaxHeap() *Heap {
	return &Heap{values: []int{}, multiplicator: 1}
}

func NewMaxHeapWithSize(size int) *Heap {
	return &Heap{values: make([]int, 0, size), multiplicator: 1}
}

func NewMinHeap() *Heap {
	return &Heap{values: []int{}, multiplicator: -1}
}

func NewMinHeapWithSize(size int) *Heap {
	return &Heap{values: make([]int, 0, size), multiplicator: -1}
}

func (h *Heap) Add(value int) {
	h.values = append(h.values, value)
	h.shiftUp(len(h.values) - 1)
}

func (h *Heap) shiftUp(itemIdx int) {
	currentIdx := itemIdx
	for currentIdx > 0 {
		parent := (currentIdx - 1) / 2
		if h.values[currentIdx]*h.multiplicator <= h.values[parent]*h.multiplicator {
			break
		}
		h.values[currentIdx], h.values[parent] = h.values[parent], h.values[currentIdx]
		currentIdx = parent
	}
}

func (h *Heap) Peak() (int, bool) {
	size := len(h.values)
	if size == 0 {
		return 0, false
	}

	return h.values[0], true
}

func (h *Heap) Pop() (int, bool) {
	size := len(h.values)
	if size == 0 {
		return 0, false
	}
	root := h.values[0]
	h.values[0] = h.values[size-1]
	h.values = h.values[:size-1]
	if len(h.values) > 0 {
		h.shiftDown(0)
	}
	return root, true
}

func (h *Heap) shiftDown(i int) {
	size := len(h.values)
	currentIdx := i
	for {
		left := 2*currentIdx + 1
		if left >= size {
			return
		}
		smallest := left
		right := left + 1
		if right < size && h.values[right]*h.multiplicator > h.values[left]*h.multiplicator {
			smallest = right
		}
		if h.values[currentIdx]*h.multiplicator >= h.values[smallest]*h.multiplicator {
			return
		}
		h.values[currentIdx], h.values[smallest] = h.values[smallest], h.values[currentIdx]
		currentIdx = smallest
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.values) == 0
}
