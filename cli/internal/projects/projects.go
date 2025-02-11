package projects

import (
	"encoding/json"
	"fmt"
	"os"
)

var projectsFileName = "./projects.json"

type Project struct {
	Name         string   `json:"name"`
	Repositories []string `json:"repositories"`
}

func saveProjectsFile(projects []Project) error {
	decodedJsonFile := struct {
		Projects []Project
	}{Projects: projects}

	jsonData, err := json.Marshal(decodedJsonFile)

	if err != nil {
		return err
	}

	os.WriteFile(projectsFileName, jsonData, 0644)

	return nil
}

func getOrCreateProjectsFile() ([]byte, error) {
	file, err := os.ReadFile(projectsFileName)

	if err != nil {
		defaultProjectsFileContent := []byte(`{"projects": []}`)
		err := os.WriteFile(projectsFileName, defaultProjectsFileContent, 0644)

		if err != nil {
			return make([]byte, 0), err
		}

		return defaultProjectsFileContent, nil
	}

	return file, nil
}

func getProjectsFromFile() ([]Project, error) {
	var decodedJsonFile struct {
		Projects []Project
	}

	file, err := getOrCreateProjectsFile()

	if err != nil {
		fmt.Println("could not read projects list")
		return make([]Project, 0), err
	}

	if err = json.Unmarshal(file, &decodedJsonFile); err != nil {
		fmt.Println("projects file is corrupted")
		return make([]Project, 0), err
	}

	return decodedJsonFile.Projects, nil
}

func ListProjects() ([]Project, error) {
	return getProjectsFromFile()
}

func AddProject(name string) (Project, error) {
	projects, err := getProjectsFromFile()

	if err != nil {
		return Project{}, err
	}

	for i := 0; i < len(projects); i++ {
		project := projects[i]

		if project.Name == name {
			return Project{}, fmt.Errorf("a project with the %v name already exists", name)
		}
	}

	newProject := Project{Name: name}
	projects = append(projects, newProject)

	if err = saveProjectsFile(projects); err != nil {
		fmt.Println("could not save projects")
		return newProject, err
	}

	return newProject, nil
}
