package webhook

import (
	"github-bot/pkg/api"
	"github-bot/pkg/config"
	"strconv"
	"strings"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// GitLabPipelineHandler handles gitlab pipeline events
func GitLabPipelineHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(gitlab.PipelineEventPayload)
	state := pl.ObjectAttributes.Status
	refs := strings.Split(pl.ObjectAttributes.Ref, "-")
	id, _ := strconv.ParseInt(refs[len(refs)-1], 10, 64)

	conf := config.LoadJobJSON().NewProject(pl.Project.Name)
	job := conf.GetProject(pl.Project.Name).GetJob(id)
	job.State = state

	stat := &api.GitHubStatus{
		State:       state,
		Context:     "gitlab-ci/pipeline",
		Description: "Pipeline " + state,
		TargetURL:   pl.Project.WebURL + "/pipelines",
	}
	stat.CheckStatus()
	config.SaveJobAsJSON(conf)
	api.CreateGitHubStatus(
		pl.Project.Namespace,
		pl.Project.Name,
		job.HeadSha,
		stat,
	)
}
