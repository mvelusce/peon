package loading_project

import "fmt"

func loadProject() {

	modules := LoadModules("testdata")

	fmt.Println(modules)
}
