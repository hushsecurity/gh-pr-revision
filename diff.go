package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/cli/cli/v2/pkg/iostreams"
)

func diffRevisions(args DiffArgs, tool string) error {
	ioStreams := iostreams.System()
	ioStreams.StartProgressIndicator()
	defer ioStreams.StopProgressIndicator()

	if args.From == args.To {
		return fmt.Errorf("cannot compare same revision: %d", args.From)
	}

	pr, err := getPullRequest()
	if err != nil {
		return err
	}
	revisions, err := parseRevisions(pr)
	if err != nil {
		return err
	}
	if len(revisions) == 0 {
		return fmt.Errorf("no revisions found")
	}

	var fromHash, toHash string
	for _, r := range revisions {
		switch r.Number {
		case args.From:
			fromHash = r.Hash
		case args.To:
			toHash = r.Hash
		}
	}
	if args.From == 0 {
		fromHash = revisions[0].BaseHash
	}
	if args.To == 0 {
		toHash = revisions[0].BaseHash
	}
	if len(fromHash) == 0 {
		return fmt.Errorf("FROM revision %d not found", args.From)
	} else if len(toHash) == 0 {
		return fmt.Errorf("TO revision %d not found", args.To)
	}

	if !hasCommit(fromHash) {
		return fmt.Errorf("FROM hash not found, try 'git fetch': %s", fromHash)
	}
	if !hasCommit(toHash) {
		return fmt.Errorf("TO hash not found, try 'git fetch': %s", toHash)
	}

	toolArgs := []string{tool}
	if tool == "difftool" {
		toolArgs = append(toolArgs, "--no-prompt")
	}
	toolArgs = append(toolArgs, fromHash)
	toolArgs = append(toolArgs, toHash)

	ioStreams.StopProgressIndicator()

	cmd := exec.CommandContext(context.Background(), "git", toolArgs...)
	cmd.Stdin = ioStreams.In
	cmd.Stdout = ioStreams.Out
	cmd.Stderr = ioStreams.ErrOut
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("'git %s' failed: %v", tool, err)
	}
	return nil
}
