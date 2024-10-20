package main

import (
	"fmt"
	"os"

	"github.com/cli/cli/v2/pkg/iostreams"
)

func listRevisions(args ListArgs) error {
	ioStreams := iostreams.System()
	ioStreams.StartProgressIndicator()
	defer ioStreams.StopProgressIndicator()

	pr, err := getPullRequest()
	if err != nil {
		return err
	}

	revisions, err := parseRevisions(pr)
	if err != nil {
		return fmt.Errorf("failed to parse revisions: %v", err)
	}

	ioStreams.StopProgressIndicator()

	for _, r := range revisions {
		r.Dump(os.Stdout, args.Verbose)
		if args.Verbose {
			fmt.Fprintf(os.Stdout, "\n")
		}
	}

	return nil
}
