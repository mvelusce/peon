package project

import (
	"container/heap"
)

type ModulePriority struct {
	module   *Module
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ModulePriority

func Init(modules []*Module) *PriorityQueue {

	priorities := map[string]int{}
	mods := map[string]*Module{}

	for _, mod := range modules {
		priorities[mod.Name] = len(mod.Dependencies)
		mods[mod.Name] = mod
	}

	pq := make(PriorityQueue, len(priorities))
	i := 0
	for key, priority := range priorities {
		pq[i] = &ModulePriority{
			module:   mods[key],
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	return &pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ModulePriority)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and module of an ModulePriority in the queue.
func (pq *PriorityQueue) update(item *ModulePriority, value *Module, priority int) {
	item.module = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
/*func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for module, priority := range items {
		pq[i] = &ModulePriority{
			module:    module,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &ModulePriority{
		module:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.module, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*ModulePriority)
		fmt.Printf("%.2d:%s ", item.priority, item.module)
	}
}*/
