package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Repository a repos struct
type Repository struct {
	Name       string `json:"name"`
	Origin     string `json:"origin"`
	OriginName string `json:"origin_name"`
	Remote     string `json:"remote"`
	RemoteName string `json:"remote_name"`
	Path       string `json:"repos_path"`
}

var (
	repos = make([]Repository, 0)
	one   sync.Once
)

const (
	reposPath    = "/etc/github-bot/config.json"
	reposPathKey = "REPO_FILE_PATH"
)

func reposFilePath() string {
	if path := os.Getenv(reposPathKey); path != "" {
		return path
	}
	return reposPath
}

// LoadRepositoryConfig load repository config
func LoadRepositoryConfig() {
	one.Do(func() {
		raw, err := ioutil.ReadFile(reposFilePath())
		if err != nil {
			log.Printf("Info: %v", err)
		}
		json.Unmarshal(raw, &repos)
	})
}

// GetRepository get repos from project name
func GetRepository(name string) *Repository {
	for i := range repos {
		if repos[i].Name == name {
			return &repos[i]
		}
	}
	log.Fatalln("Can't not find repository.")
	return nil
}
