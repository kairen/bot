package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Job a job struct
type Job struct {
	ID      int64  `json:"id"`
	HeadSha string `json:"head_sha"`
	State   string `json:"state"`
}

// Project a project struct
type Project struct {
	Name string `json:"name"`
	Jobs []Job  `json:"jobs"`
}

// ProjectList a project list
type ProjectList []Project

const (
	jobPath    = "/var/lib/github-bot/jobs.json"
	jobPathKey = "JOB_FILE_PATH"
	// StateDefault is state default value
	StateDefault = "created"
)

// GetJob add new job
func (p *Project) GetJob(id int64) *Job {
	for i := range p.Jobs {
		if p.Jobs[i].ID == id {
			return &p.Jobs[i]
		}
	}
	return nil
}

// AddJob add new job
func (p *Project) AddJob(id int64) {
	t := true
	for i := range p.Jobs {
		if p.Jobs[i].ID == id {
			t = false
			break
		}
	}
	if t {
		j := Job{ID: id, State: StateDefault}
		p.Jobs = append(p.Jobs, j)
	}
}

// RemoveJob remote a old job
func (p *Project) RemoveJob(id int64) bool {
	for i := range p.Jobs {
		if p.Jobs[i].ID == id {
			p.Jobs = p.Jobs[:i+copy(p.Jobs[i:], p.Jobs[i+1:])]
			return true
		}
	}
	return false
}

// NewProject get project from name
func (pl ProjectList) NewProject(name string) ProjectList {
	t := true
	for i := range pl {
		if pl[i].Name == name {
			t = false
			break
		}
	}
	if t {
		pl = append(pl, Project{Name: name})
	}
	return pl
}

// GetProject get project from name
func (pl ProjectList) GetProject(name string) *Project {
	for i := range pl {
		if pl[i].Name == name {
			return &pl[i]
		}
	}
	return nil
}

func jobFilePath() string {
	if path := os.Getenv(jobPathKey); path != "" {
		return path
	}
	return jobPath
}

// LoadJobJSON list a json data
func LoadJobJSON() ProjectList {
	raw, err := ioutil.ReadFile(jobFilePath())
	if err != nil {
		log.Printf("Info: %v", err)
		raw = []byte(`[]`)
	}

	var pl ProjectList
	json.Unmarshal(raw, &pl)
	return pl
}

// SaveJobAsJSON save state to json file
func SaveJobAsJSON(v interface{}) {
	fo, err := os.Create(jobFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	e := json.NewEncoder(fo)
	if err := e.Encode(v); err != nil {
		log.Fatal(err)
	}
}
