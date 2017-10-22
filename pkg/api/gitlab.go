package api

import (
	"log"
	"sync"

	gitlab "github.com/xanzy/go-gitlab"
)

// A GitLabInfo is gitlab account info
type GitLabInfo struct {
	token    string
	endpoint string
}

var (
	lc *gitlab.Client
	lo sync.Once
)

// NewGitLab returns a github using the given endpoint and token
func NewGitLab(token string, ep string) *GitLabInfo {
	return &GitLabInfo{token: token, endpoint: ep}
}

// InitGitLabClient init gitlab client
func (inf *GitLabInfo) InitGitLabClient() *gitlab.Client {
	lo.Do(func() {
		lc = gitlab.NewClient(nil, inf.token)
		lc.SetBaseURL(inf.endpoint + "/api/v4")
		if lc == nil {
			log.Fatal("GitLab client init  falied...")
		}
	})
	return lc
}

// CreateGitLabPRComment comment a message
func CreateGitLabPRComment(owner string, repo string, number int, body string) {
	opts := &gitlab.CreateMergeRequestNoteOptions{
		Body: gitlab.String(body),
	}
	pid := owner + "/" + repo
	_, _, err := lc.Notes.CreateMergeRequestNote(pid, int(number), opts)
	if err != nil {
		log.Fatal(err)
	}
}
