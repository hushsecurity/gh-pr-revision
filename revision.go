package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	mdPrefix  = "[//]: # (gh-revision:"
	mdSuffix  = ")"
	bodyStart = "[//]: # (gh-revision-body-start)"
	bodyEnd   = "[//]: # (gh-revision-body-end)"
)

type Revision struct {
	CreatedAt string `json:"-"`
	Number    uint64 `json:"number"`
	Hash      string `json:"hash"`
	BaseHash  string `json:"baseHash"`
	Comment   string `json:"-"`
}

func (r Revision) Dump(w io.Writer, long bool) {
	if long {
		fmt.Fprintf(w, "%d. %s\n", r.Number, r.CreatedAt)
		fmt.Fprintf(w, "Hash: %s\n", r.Hash)
		fmt.Fprintf(w, "BaseHash: %s\n", r.BaseHash)
		if len(r.Comment) > 0 {
			fmt.Fprintf(w, "Comment:\n%s\n", r.Comment)
		}
	} else {
		fmt.Fprintf(w, "%2d. %s %s\n", r.Number, r.CreatedAt, r.Hash)
	}
}

func parseRevisions(pr PullRequest) (revisions []Revision, err error) {
	var rev *Revision
	for _, comment := range pr.Comments {
		if rev, err = parseRevision(comment); err != nil {
			return nil, err
		}
		if rev != nil {
			revisions = append(revisions, *rev)
		}
	}
	return revisions, nil
}

func parseRevision(comment Comment) (*Revision, error) {
	for _, s := range strings.Split(comment.Body, "\n") {
		if tmp, ok := strings.CutPrefix(s, mdPrefix); ok {
			if tmp, ok := strings.CutSuffix(tmp, mdSuffix); ok {
				return parseRevisionStr(comment, tmp)
			}
		}
	}
	return nil, nil
}

func parseRevisionStr(comment Comment, encoded string) (*Revision, error) {
	var rev Revision
	if err := json.Unmarshal([]byte(encoded), &rev); err != nil {
		return nil, fmt.Errorf("json.decode failed: %v", err)
	}
	rev.CreatedAt = comment.CreatedAt
	rev.Comment = parseBody(comment)
	return &rev, nil
}

func encodeRevision(rev Revision) (string, error) {
	json, err := json.Marshal(rev)
	if err != nil {
		return "", fmt.Errorf("json.encode failed: %v", err)
	}
	return fmt.Sprintf("%s%s%s", mdPrefix, string(json), mdSuffix), nil
}

func encodeBody(body string) string {
	return fmt.Sprintf("\n\n%s\n\n%s\n\n%s\n\n", bodyStart, body, bodyEnd)
}

func parseBody(comment Comment) string {
	var startSeen, endSeen bool
	var lines []string

	for _, line := range strings.Split(comment.Body, "\n") {
		if !startSeen && line == bodyStart {
			startSeen = true
			continue
		}
		if !startSeen {
			continue
		}
		if line == bodyEnd {
			endSeen = true
			break
		}
		lines = append(lines, line)
	}

	if !(startSeen && endSeen) {
		return ""
	}

	// strip empty lines in the beginning
	for len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}

	// strip empty lines in the end
	for len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return strings.Join(lines, "\n")
}
