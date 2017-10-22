[![Docker Build Statu](https://img.shields.io/docker/build/kairen/github-bot.svg)](https://hub.docker.com/r/kairen/github-bot/) [![Build Status](https://travis-ci.org/kairen/github-bot.svg?branch=master)](https://travis-ci.org/kairen/github-bot) [![Coverage Status](https://coveralls.io/repos/github/kairen/github-bot/badge.svg?branch=master)](https://coveralls.io/github/kairen/github-bot?branch=master)
# Webhook handler for KaiRen's Bot
Make GitLab CI work with a GitHub repository written in Go. The bot watch webhook event to handle something.

![snapshot](logo.png)

## Build
To build bot into a container via Docker:
```sh
$ docker build -t kairen/bot:0.1.0 .
```

## Usage
Run bot on Docker as below command:
```sh
$ docker run --rm -ti -e [envs] kairen/bot:0.1.0 [options]
$ docker run --rm -ti kairen/bot:0.1.0
Usage: bot [OPTION]...
bot watch your GitHub and GitLab event to handle anything!!

Options:
  -github-token string
    	GitHub access token.
  -gitlab-endpoint string
    	GitLab API Endpoint. (default "https://gitlab.com/api/v4")
  -gitlab-token string
    	GitLab access token.
  -path string
    	Webhook uri path. (default "/webhooks")
  -port int
    	Webhook server port. (default 8080)
  -secret string
    	Webhook server secret. (default "6324d8bf")
```
