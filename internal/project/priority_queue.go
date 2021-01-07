package project

import (
	"container/heap"

	log "github.com/sirupsen/logrus"
)

// A ModulePriority holds details of the module and its priority
type ModulePriority struct {
	module   *Module
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds modulePriority.
type PriorityQueue struct {
	priorities       []*ModulePriority
	prioritiesByName map[string]*ModulePriority
}

// Init inistialize the priority Queue
func Init(modules []*Module) *PriorityQueue {

	modDeps := map[string]int{}
	mods := map[string]*Module{}

	for _, mod := range modules {
		modDeps[mod.Name] = len(mod.Dependencies)
		mods[mod.Name] = mod
	}

	priorities := make([]*ModulePriority, len(modDeps))
	prioritiesByName := map[string]*ModulePriority{}
	i := 0
	for key, priority := range modDeps {
		modPriority := &ModulePriority{
			module:   mods[key],
			priority: priority,
			index:    i,
		}
		priorities[i] = modPriority
		prioritiesByName[key] = modPriority
		i++
	}
	pq := &PriorityQueue{
		priorities:       priorities,
		prioritiesByName: prioritiesByName,
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

// Push push object to priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.priorities)
	item := x.(*ModulePriority)
	item.index = n
	pq.priorities = append(pq.priorities, item)
}

// Pop pop object to priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.priorities
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.priorities = old[0 : n-1]
	return item
}

// Update update the priority of the object matching the key
func (pq *PriorityQueue) Update(key string, priority int) {
	modulePriority := pq.prioritiesByName[key]
	log.Debugf("Updating module: %s. Priority: %d", modulePriority.module.Name, key)
	pq.update(modulePriority, modulePriority.module, priority)
}

// Decrease decrease the priority of the object matching the key by one
func (pq *PriorityQueue) Decrease(key string) {
	modulePriority := pq.prioritiesByName[key]
	priority := modulePriority.priority - 1
	println(priority)
	log.Debugf("Decreasing module priority: %v", modulePriority.module.Name)
	pq.Update(key, priority)
}

func (pq *PriorityQueue) update(item *ModulePriority, value *Module, priority int) {
	item.module = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
