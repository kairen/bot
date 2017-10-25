package main

import (
	"errors"
	"flag"
	"fmt"
	"github-bot/pkg/config"
	"github-bot/pkg/webhook"
	"os"
)

var (
	// Error message define
	errRequired = errors.New("[ERROR] Missing required opts")

	// GitHub and GitLab opts
	githubToken    = flag.String("github-token", "", "GitHub access token.")
	gitlabToken    = flag.String("gitlab-token", "", "GitLab access token.")
	gitlabEndpoint = flag.String("gitlab-endpoint", "https://gitlab.com", "GitLab API Endpoint.")

	// Webhook server opts
	webhookPort   = flag.Int("port", 8080, "Webhook server port.")
	webhookSecret = flag.String("secret", "6324d8bf", "Webhook server secret.")
	webhookPath   = flag.String("path", "/webhooks", "Webhook uri path.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: github-bot [OPTION]...\n")
	fmt.Fprintf(os.Stderr, "github-bot watch your GitHub and GitLab event to handle anything!!\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
}

func fail(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *githubToken == "" || *gitlabToken == "" {
		usage()
		fail(errRequired)
	}

	config.LoadRepositoryConfig()

	// Init GitHub and GitLab account
	c := webhook.NewAccount(*githubToken, *gitlabToken, *gitlabEndpoint)
	c.InitAccount()

	// Start webhook server
	s := webhook.NewServer(*webhookPort, *webhookSecret, *webhookPath)
	s.RunServer()
}
