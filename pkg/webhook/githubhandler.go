package webhook

import (
	"bot/pkg/api"
	"bot/pkg/utils"
	"log"
	"strings"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

// GitHubIssueCommentHandler handles gitHub pull request events
func GitHubIssueCommentHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(github.IssueCommentPayload)
	if pl.Action == "created" && isPullRequestComment(pl.Issue.HTMLURL) {
		switch pl.Comment.Body {
		case "/ok-to-test":
			owner := isOwner(&pl)
			if owner {
				path := "/tmp/" + pl.Repository.Name
				remoteURL := account.gitlabEndpoint + "/" + pl.Repository.FullName
				utils.GitClone(pl.Repository.CloneURL, path)
				utils.GitAddRemote(path, "gitlab", remoteURL)
				utils.GitFetchLastUpdate(path, "origin")
			}
		default:
			log.Print("Other Event trigger")
		}
	}
}

// GitHubPullRequestHandler handles github pull request events
func GitHubPullRequestHandler(payload interface{}, header webhooks.Header) {
	// pl := payload.(github.PullRequestPayload)
	// pl.PullRequest.Commits
	// urls := strings.Split(pl.PullRequest.StatusesURL, "/")
	// repo := pl.Repository
}

// Check sender is owner
func isOwner(pl *github.IssueCommentPayload) bool {
	if pl.Comment.AuthorAssociation != "OWNER" {
		msg := "@" + pl.Sender.Login + ": You can't run testing, because you are not a `member`."
		repo := pl.Repository
		api.CreateGitHubPRComment(repo.Owner.Login, repo.Name, int(pl.Issue.Number), msg)
		return false
	}
	return true
}

func isPullRequestComment(url string) bool {
	i := strings.Index(url, "pull")
	return (i > 0)
}
