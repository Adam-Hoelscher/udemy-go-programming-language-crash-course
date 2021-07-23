package priority

type IntMaxHeap []int

func (h IntMaxHeap) Len() int { return len(h) }

// The heap package implements a MinHeap, which prioritizes lower values.
// Because we want to prioritize higher values, we'll say that a higher value
// is "Less" than a lower value
func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
