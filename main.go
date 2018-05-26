/*
ProjectSource Server
Version Alpha 0.0.1a


*/

package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

/* TEST DATA SET */

type Project struct {
	Name		string		`json:"name,omitempy"`
	Category	*Category	`json:"category,omitempty`

}

type Category struct {
	Title		string		`json:"title,omitempty"`

}

var projects []Project


/* ENDPOINT DESCRIPTIONS: */

/* GetProjects - GET
	- Retreive all projects from the server */
func getProjects(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(projects)
}

/* GetProject - GET
	- Retreive a single project from the server */
func getProject(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    for _, item := range projects {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Project{})
}

/* CreateProject - POST
	- Post a newly created project to the server */	
func createProject(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    var project Project
    _ = json.NewDecoder(r.Body).Decode(&project)
    project.Name = params["name"]
    projects = append(projects, project)
    json.NewEncoder(w).Encode(projects)
}

/* deleteProject - GET
	- Delete a project from the server */	
func deleteProject(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    for index, item := range projects {
        if item.Name == params["name"] {
            projects = append(projects[:index], projects[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(projects)
    }
}

// our main function
func main() {
    router := mux.NewRouter()

    //manually add some data:
	projects = append(projects, Project{Name: "ProLiant POS", Category: &Category{Title: "Electronics"}})
	projects = append(projects, Project{Name: "Michael's Grad Cap", Category: &Category{Title: "Electronics"}})
	projects = append(projects, Project{Name: "Peter's Drift Trike", Category: &Category{Title: "Hardware"}})
	projects = append(projects, Project{Name: "Toastmatic 3000", Category: &Category{Title: "Food Hacking"}})

    //Insert Endpoints here
    router.HandleFunc("/project", getProjects).Methods("GET")
    router.HandleFunc("/project/{id}", getProject).Methods("GET")
    router.HandleFunc("/project/{id}", createProject).Methods("POST")
    router.HandleFunc("/project/{id}", deleteProject).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}