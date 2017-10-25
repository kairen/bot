package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func logIfError(err error) {
	if err != nil {
		log.Printf("INFO: %v", err)
	}
}

func git(path, argstr string) {
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal(err)
	}
	head := "git"
	if path != "" {
		argstr = fmt.Sprintf("-C %s ", path) + argstr
	}
	args := splitArgs(argstr)
	_, execErr := exec.Command(head, args...).Output()
	logIfError(execErr)
}

func splitArgs(argstr string) []string {
	args := strings.Split(argstr, " ")
	return args
}

// GitClone clone project from url
func GitClone(path, url string) {
	log.Println("Git remote clone")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		argstr := fmt.Sprintf("clone %s %s", url, path)
		git("", argstr)
	}
}

// GitAddRemote add remote into project
func GitAddRemote(path, remote, url string) {
	log.Println("Git remote add")
	argstr := fmt.Sprintf("remote add %s %s", remote, url)
	git(path, argstr)
}

// GitFetch fetch remote changes
func GitFetch(path, remote string, pid int64) {
	log.Println("Git fetch")
	argstr := fmt.Sprintf("fetch %s refs/pull/%d/head:pr-%d", remote, pid, pid)
	git(path, argstr)
}

// GitPushAndDelete push and delete branch
func GitPushAndDelete(path, remote string, pid int64) {
	log.Println("Git push")
	argstr := fmt.Sprintf("push %s pr-%d", remote, pid)
	git(path, argstr)

	log.Println("Git branch delete")
	argstr = fmt.Sprintf("branch -D pr-%d", pid)
	git(path, argstr)
}
