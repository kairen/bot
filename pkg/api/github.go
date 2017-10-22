package api

import (
	"context"
	"log"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// A GitHubInfo is github account info
type GitHubInfo struct {
	token string
}

var (
	hc *github.Client
	ho sync.Once
)

// NewGitHub returns a github using the given token
func NewGitHub(token string) *GitHubInfo {
	return &GitHubInfo{token: token}
}

// InitGitHubClient github account
func (inf *GitHubInfo) InitGitHubClient() *github.Client {
	ho.Do(func() {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: inf.token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		hc = github.NewClient(tc)
		if hc == nil {
			log.Fatal("GitHub client init  falied...")
		}
	})
	return hc
}

// CreateGitHubPRComment create a pull reauest comment message
func CreateGitHubPRComment(owner string, repo string, number int, body string) {
	issue := &github.IssueComment{Body: &body}
	_, _, err := hc.Issues.CreateComment(context.Background(), owner, repo, number, issue)
	if err != nil {
		log.Fatal(err)
	}
}

// GitHubStatus a status data
type GitHubStatus struct {
	State       string
	TargetURL   string
	Context     string
	Description string
}

// CreateGitHubStatus create a status
func CreateGitHubStatus(owner string, repo string, sha string, stat *GitHubStatus) {
	status := &github.RepoStatus{
		State:       github.String(stat.State),
		TargetURL:   github.String(stat.TargetURL),
		Context:     github.String(stat.Context),
		Description: github.String(stat.Description),
	}
	_, _, err := hc.Repositories.CreateStatus(context.Background(), owner, repo, sha, status)
	if err != nil {
		log.Fatal(err)
	}
}
