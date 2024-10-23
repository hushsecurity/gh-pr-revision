#!/bin/bash -e

DIR=$1
MALFORMATTED_FILES=$(gofmt -l "$DIR")
if [[ -z "$MALFORMATTED_FILES" ]]; then
    exit 0
fi
printf "Malformatted files found:\n%s\n\n" "$MALFORMATTED_FILES"
gofmt -d "$DIR"
exit 1
