package main

import (
	"fmt"

	"github.com/EugenVolosciuc/git-report/cli/internal/projects"
)

func main() {
	// if _, err := projects.AddProject("test-8"); err != nil {
	// 	fmt.Println(err)
	// }

	if _, err := projects.DeleteProject("git-report"); err != nil {
		fmt.Println(err)
	}

	projectList, err := projects.ListProjects()

	if err != nil {
		panic(err)
	}

	fmt.Printf("projects: %+v", projectList)
}
