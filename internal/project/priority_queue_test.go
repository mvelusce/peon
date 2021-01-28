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
	pq.Decrease("mod1")
	pq.Decrease("mod2")
	assert.Equal(t, "mod1", heap.Pop(pq).(*ModulePriority).module.Name)
	pq.Decrease("mod2")
	pq.Decrease("mod4")
	assert.Equal(t, "mod2", heap.Pop(pq).(*ModulePriority).module.Name)
	pq.Decrease("mod4")
	pq.Decrease("mod3")
	assert.Equal(t, "mod4", heap.Pop(pq).(*ModulePriority).module.Name)
	assert.Equal(t, "mod3", heap.Pop(pq).(*ModulePriority).module.Name)
	pq.Decrease("mod5")
	assert.Equal(t, "mod5", heap.Pop(pq).(*ModulePriority).module.Name)
	assert.Equal(t, 0, pq.Len())
}
