package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2"
	api "github.com/cli/go-gh/v2/pkg/api"
	"github.com/kballard/go-shellquote"
)

var (
	bom = []byte{0xEF, 0xBB, 0xBF}
)

func getPullRequest() (pr PullRequest, err error) {
	stdout, stderr, err := gh.Exec("pr", "view", "--json",
		"id,number,state,isDraft,commits,comments,reviewRequests,headRepository,headRepositoryOwner")
	if err != nil {
		return pr, fmt.Errorf("'gh pr view' failed: %s", stderr.String())
	}
	if err = json.Unmarshal(stdout.Bytes(), &pr); err != nil {
		return pr, fmt.Errorf("failed to parse pr: %v", err)
	}
	return pr, nil
}

func getPullRequestFromApi(owner, repo string, number uint64) (pr ApiPullRequest, err error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return pr, fmt.Errorf("failed to create REST client: %v", err)
	}
	if err = client.Get(fmt.Sprintf("repos/%s/%s/pulls/%d", owner, repo, number), &pr); err != nil {
		return pr, fmt.Errorf("failed to get pull request from api: %v", err)
	}
	return pr, nil
}

func getRepository() (repo Repository, err error) {
	stdout, stderr, err := gh.Exec("repo", "view", "--json", "url")
	if err != nil {
		return repo, fmt.Errorf("'gh repo view' failed: %s", stderr.String())
	}
	if err = json.Unmarshal(stdout.Bytes(), &repo); err != nil {
		return repo, fmt.Errorf("failed to parse repo: %v", err)
	}
	return repo, nil
}

func getDefaultEditor() string {
	if e := os.Getenv("GIT_EDITOR"); e != "" {
		return e
	} else if e := os.Getenv("VISUAL"); e != "" {
		return e
	} else if e := os.Getenv("EDITOR"); e != "" {
		return e
	} else if runtime.GOOS == "windows" {
		return "notepad"
	} else {
		return "nano"
	}
}

func getEditor() (string, error) {
	stdout, _, err := gh.Exec("config", "get", "editor")
	if err != nil {
		return "", fmt.Errorf("'gh config get editor' failed: %v", err)
	}
	if editor := strings.TrimSpace(stdout.String()); len(editor) > 0 {
		return editor, nil
	}
	return getDefaultEditor(), nil
}

func revParse(ref string) (string, error) {
	out, err := exec.Command("git", "rev-parse", ref).Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("'git rev-parse %s' failed: %v\nstdout=\n%s\nstderr=\n%s",
				ref, ee, string(out), string(ee.Stderr))
		}
		return "", fmt.Errorf("'git rev-parse %s' failed: %v\nstdout=\n%s", ref, err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

func addPrComment(path string) error {
	_, _, err := gh.Exec("pr", "comment", "-F", path)
	if err != nil {
		return fmt.Errorf("'gh pr comment' failed: %v", err)
	}
	return nil
}

func editUserComment(ioStreams *iostreams.IOStreams) (string, error) {
	tmpFile, err := os.CreateTemp("", "gh-pr-revision-user-comment.*.md")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if runtime.GOOS == "windows" {
		if _, err := tmpFile.Write(bom); err != nil {
			return "", fmt.Errorf("failed to write bom: %v", err)
		}
	}

	if err := tmpFile.Close(); err != nil {
		return "", fmt.Errorf("failed to close tmp file: %v", err)
	}

	editor, err := getEditor()
	if err != nil {
		return "", fmt.Errorf("failed to determine editor: %v", err)
	}

	args, err := shellquote.Split(editor)
	if err != nil {
		return "", fmt.Errorf("shellquote.Split failed: %s", err)
	}

	args = append(args, tmpFile.Name())

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = ioStreams.In
	cmd.Stdout = ioStreams.Out
	cmd.Stderr = ioStreams.ErrOut
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("editor '%s' failed: %v", args[0], err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to read edited message: %v", err)
	}

	return string(bytes.TrimPrefix(data, bom)), nil
}
