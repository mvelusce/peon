package project

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	var modules = []*Module{
		{"mod3", "mod3", []string{"mod2"}},
		{"mod5", "mod5", []string{"mod3"}},
		{"mod0", "mod0", []string{}},
		{"mod1", "mod1", []string{"mod0"}},
		{"mod4", "mod4", []string{"mod2", "mod1"}},
		{"mod2", "mod2", []string{"mod1", "mod0"}},
	}

	pq := Init(modules)

	assert.Equal(t, 6, pq.Len())
	assert.Equal(t, "mod0", heap.Pop(pq).(*ModulePriority).module.Name)
	assert.Equal(t, 5, pq.Len())
	//pq.Decrease("mod1")
	println(heap.Pop(pq).(*ModulePriority).module.Name)
	println(heap.Pop(pq).(*ModulePriority).module.Name)
	println(heap.Pop(pq).(*ModulePriority).module.Name)
	println(heap.Pop(pq).(*ModulePriority).module.Name)
	println(heap.Pop(pq).(*ModulePriority).module.Name)
	//assert.Equal(t, "mod1", heap.Pop(pq).(*ModulePriority).module.Name)
}
