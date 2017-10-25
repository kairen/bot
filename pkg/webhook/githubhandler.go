package webhook

import (
	"github-bot/pkg/api"
	"github-bot/pkg/config"
	"github-bot/pkg/utils"
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
				repo := config.GetRepository(pl.Repository.Name)
				utils.GitClone(repo.Path, pl.Repository.CloneURL)
				utils.GitAddRemote(repo.Path, repo.RemoteName, repo.Remote)
				utils.GitFetch(repo.Path, repo.OriginName, pl.Issue.Number)
				utils.GitPushAndDelete(repo.Path, repo.RemoteName, pl.Issue.Number)
			}
		default:
			log.Print("Other Event trigger")
		}
	}
}

// GitHubPullRequestHandler handles github pull request events
func GitHubPullRequestHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PullRequestPayload)
	conf := config.LoadJobJSON().NewProject(pl.Repository.Name)
	project := conf.GetProject(pl.Repository.Name)

	switch pl.Action {
	case "opened", "edited":
		project.AddJob(pl.PullRequest.Number)
		job := project.GetJob(pl.PullRequest.Number)
		job.HeadSha = pl.PullRequest.Head.Sha
	case "closed":
		if pl.PullRequest.Merged {
			project.RemoveJob(pl.PullRequest.Number)
		}
	}
	config.SaveJobAsJSON(conf)
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
