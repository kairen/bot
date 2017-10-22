package webhook

import (
	"log"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// GitLabPipelineHandler handles gitlab pipeline events
func GitLabPipelineHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(gitlab.PipelineEventPayload)
	log.Printf("%+v", pl.Project.Name)
}
