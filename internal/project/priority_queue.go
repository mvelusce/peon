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
type PriorityQueue struct {
	priorities    []*ModulePriority
	moduleIndexes map[string]int
}

func Init(modules []*Module) *PriorityQueue {

	modDeps := map[string]int{}
	mods := map[string]*Module{}

	for _, mod := range modules {
		modDeps[mod.Name] = len(mod.Dependencies)
		mods[mod.Name] = mod
	}

	priorities := make([]*ModulePriority, len(modDeps))
	moduleIndexes := map[string]int{}
	i := 0
	for key, priority := range modDeps {
		priorities[i] = &ModulePriority{
			module:   mods[key],
			priority: priority,
			index:    i,
		}
		moduleIndexes[key] = i
		i++
	}
	pq := &PriorityQueue{
		priorities:    priorities,
		moduleIndexes: moduleIndexes,
	}
	heap.Init(pq)
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq.priorities) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.priorities[i].priority < pq.priorities[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.priorities[i], pq.priorities[j] = pq.priorities[j], pq.priorities[i]
	pq.priorities[i].index = i
	pq.priorities[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.priorities)
	item := x.(*ModulePriority)
	item.index = n
	pq.priorities = append(pq.priorities, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.priorities
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.priorities = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(key string, priority int) {
	moduleIndex := pq.moduleIndexes[key]
	modulePriority := pq.priorities[moduleIndex]
	pq.update(modulePriority, modulePriority.module, priority)
}

func (pq *PriorityQueue) Decrease(key string) {
	moduleIndex := pq.moduleIndexes[key]
	modulePriority := pq.priorities[moduleIndex]
	priority := modulePriority.priority - 1
	println(priority)
	pq.Update(key, priority)
}

func (pq *PriorityQueue) update(item *ModulePriority, value *Module, priority int) {
	item.module = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
