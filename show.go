package main

import (
	"fmt"
	"os"

	"github.com/cli/cli/v2/pkg/iostreams"
)

func showRevision(args ShowArgs) error {
	ioStreams := iostreams.System()
	ioStreams.StartProgressIndicator()
	defer ioStreams.StopProgressIndicator()

	pr, err := getPullRequest()
	if err != nil {
		return err
	}

	revisions, err := parseRevisions(pr)
	if err != nil {
		return err
	}

	ioStreams.StopProgressIndicator()

	for _, r := range revisions {
		if r.Number == args.Number {
			r.Dump(os.Stdout, true)
			return nil
		}
	}

	return fmt.Errorf("revision %d not found", args.Number)
}
