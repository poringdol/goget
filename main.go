package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	src = "src"

	httpsPrefix = "https://"
	gitPrefix   = "git@"
	gitSuffix   = ".git"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a repository")
	}

	link := os.Args[1]
	if link == "" {
		log.Fatal("Please provide a repository")
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("Please set GOPATH environment variable")
	}

	dstPath := getDstPath(link, gopath)

	cmd := exec.Command("git", "clone", link, dstPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error cloning repository %s: %s", link, err)
	}
}

func getDstPath(link, gopath string) string {
	link = strings.TrimSuffix(link, gitSuffix)

	if strings.HasPrefix(link, httpsPrefix) {
		link = strings.TrimPrefix(link, httpsPrefix)
	}

	if strings.HasPrefix(link, gitPrefix) {
		link = strings.TrimPrefix(link, gitPrefix)
		link = strings.Replace(link, ":", "/", 1)
	}

	return filepath.Join(gopath, src, link)
}
