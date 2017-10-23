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
	if state != "running" && state != "canceled" {
		refs := strings.Split(pl.ObjectAttributes.Ref, "-")
		conf := config.LoadJobJSON().NewProject(pl.Project.Name)
		project := conf.GetProject(pl.Project.Name)
		id, _ := strconv.ParseInt(refs[len(refs)-1], 10, 64)
		job := project.GetJob(id)
		job.State = state
		config.SaveJobAsJSON(conf)
		api.CreateGitHubStatus(
			pl.Project.Namespace,
			pl.Project.Name,
			job.HeadSha,
			&api.GitHubStatus{
				State:       state,
				Context:     "gitlab-ci/pipeline",
				Description: "Pipeline " + state,
				TargetURL:   pl.Project.WebURL,
			},
		)
	}
}
