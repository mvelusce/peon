package project

import (
	"testing"
)

func TestLoadingProject(t *testing.T) {

	project := LoadProject()

	/*for v := 0; v < project.dependencies.Order(); v++ {
		fmt.Println(v)
		fmt.Println("Install: ", project.modules[v].Name)
	}*/

	//project.Build()

	project.BuildModule("module_d")
}
