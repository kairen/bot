package utils

import (
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func logIfError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func repos(path string) *git.Repository {
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// GitClone clone project from url
func GitClone(url, path string) {
	_, ferr := os.Stat(path)
	if os.IsNotExist(ferr) {
		opts := &git.CloneOptions{URL: url, Progress: os.Stdout}
		_, err := git.PlainClone(path, false, opts)
		logIfError(err)
	}
}

// GitPull pull repos to update changes
func GitPull(path, remoteName, ref string) {
	r := repos(path)
	w, err := r.Worktree()
	logIfError(err)

	err = w.Pull(&git.PullOptions{
		RemoteName:    remoteName,
		ReferenceName: plumbing.ReferenceName(ref),
	})
	logIfError(err)
}

// GitAddRemote add remote into project
func GitAddRemote(path, remoteName, remoteURL string) {
	r := repos(path)
	_, err := r.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{remoteURL},
	})
	logIfError(err)
}

// GitFetchLastUpdate fetch remote changes
func GitFetchLastUpdate(path, remoteName string) {
	r := repos(path)
	ref := config.RefSpec("refs/pull/*/head:refs/remotes/pr-*")
	err := r.Fetch(&git.FetchOptions{
		RemoteName: remoteName,
		RefSpecs:   []config.RefSpec{ref},
	})
	logIfError(err)
}
