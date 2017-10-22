package webhook

import (
	"bot/pkg/api"
	"log"
	"strconv"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// Account is GitHub and GitLab info
type Account struct {
	githubToken    string
	gitlabToken    string
	gitlabEndpoint string
}

// Server is webhook server
type Server struct {
	port   int
	secret string
	path   string
}

var account *Account

// NewAccount returns a client using the given github token,.., etc
func NewAccount(ht, lt, ep string) *Account {
	account = &Account{githubToken: ht, gitlabToken: lt, gitlabEndpoint: ep}
	return account
}

// InitAccount is init github and gitlab account
func (a *Account) InitAccount() {
	ghc := api.NewGitHub(a.githubToken)
	ghc.InitGitHubClient()

	glc := api.NewGitLab(a.gitlabToken, a.gitlabEndpoint)
	glc.InitGitLabClient()
}

// NewServer returns a server using the given port,.., etc
func NewServer(port int, secret, path string) *Server {
	return &Server{port: port, secret: secret, path: path}
}

// RunServer is run webhook server
func (s *Server) RunServer() {
	gh := github.New(&github.Config{Secret: s.secret})
	gh.RegisterEvents(GitHubIssueCommentHandler, github.IssueCommentEvent)
	gh.RegisterEvents(GitHubPullRequestHandler, github.PullRequestEvent)

	gl := gitlab.New(&gitlab.Config{Secret: s.secret})
	gl.RegisterEvents(GitLabPipelineHandler, gitlab.PipelineEvents)

	hooks := []webhooks.Webhook{gh, gl}
	err := webhooks.Run(hooks, ":"+strconv.Itoa(s.port), s.path)
	if err != nil {
		log.Print(err)
	}
}
