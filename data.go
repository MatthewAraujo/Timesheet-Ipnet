package main

import (
	"strconv"

	"github.com/google/uuid"
)

var dataClient []Client
var dataProject []Project

type Project struct {
	ID        string
	Project   string
	Alias     string
	Client_Id string
}

type Client struct {
	ID     string
	Client string
	Alias  string
}

func init() {
	dataClient = []Client{
		{
			ID:     uuid.New().String(),
			Client: "Ipnet Growth Parter",
			Alias:  "IPNET",
		},
		{
			ID:     uuid.New().String(),
			Client: "Blue One Express Cargo Logistica",
			Alias:  "Blue Express",
		},
		{
			ID:     uuid.New().String(),
			Client: "ABC BRASIL",
			Alias:  "ABC",
		},
	}

	dataProject = []Project{
		{
			ID:        uuid.New().String(),
			Project:   "WIP Timesheet Calendar",
			Alias:     "P2401",
			Client_Id: dataClient[0].ID,
		},
	}
}

func getProjectByID(id string) Project {
	var result Project
	for _, i := range dataProject {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func updateProject(project Project) {
	result := []Project{}
	for _, i := range dataProject {
		if i.ID == project.ID {
			i.Project = project.Project
			i.Alias = project.Alias
			i.Client_Id = project.Client_Id
		}
		result = append(result, i)
	}
	dataProject = result

}

func addProject(project Project) {
	max := 0
	for _, i := range dataProject {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	dataProject = append(dataProject, Project{
		ID:        id,
		Project:   project.Project,
		Alias:     project.Alias,
		Client_Id: project.Client_Id,
	})

}

func deleteProject(id string) {
	result := []Project{}
	for _, i := range dataProject {
		if i.ID != id {
			result = append(result, i)
		}
	}
	dataProject = result
}

func getClientByID(id string) Client {
	var result Client
	for _, i := range dataClient {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func updateClient(client Client) {
	result := []Client{}
	for _, i := range dataClient {
		if i.ID == client.ID {
			i.Client = client.Client
			i.Alias = client.Alias
		}
		result = append(result, i)
	}
	dataClient = result
}

func addClient(client Client) {
	max := 0
	for _, i := range dataClient {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	dataClient = append(dataClient, Client{
		ID:     id,
		Client: client.Client,
		Alias:  client.Alias,
	})
}

func deleteClient(id string) {
	result := []Client{}
	for _, i := range dataClient {
		if i.ID != id {
			result = append(result, i)
		}
	}
	dataClient = result
}
