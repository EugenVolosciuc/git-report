package main

import (
	"fmt"

	"github.com/EugenVolosciuc/git-report/cli/internal/projects"
)

func main() {
	if _, err := projects.AddProject("git-report"); err != nil {
		fmt.Println(err)
	}

	projectList, err := projects.ListProjects()

	if err != nil {
		panic(err)
	}

	fmt.Println("projects: ", projectList)
}
