package loading_project

import (
	"fmt"
	"testing"
)

func TestLoadingProject(t *testing.T) {

	modules, g := loadProject()

	for v := 0; v < g.Order(); v++ {
		fmt.Println(v)
		fmt.Println("Install: ", modules[v].Name)
	}
}
