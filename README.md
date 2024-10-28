# gh-pr-revision

A GitHub CLI extension for managing pull request revisions.

## Motivation

When we started using GitHub pull requests for code review one thing was immediately
missing relative to our old review practice, a review revision.
Once a pull request is created in its initial version all subsequent versions appear on
its timeline as force-pushes. A developer may make several force-pushes
to fix comments of one review. This makes the reviewer's job of tracking **incremental**
changes harder.

To make incremental changes more visible and easier to track we decided to add pull
request comments which summarize what was changed since the last review.
To automate this process and add diff links to the previous revision the
idea of `gh-pr-revision` CLI extension was born.

## Installation

To install the latest version of the extension use the following command:

```shell
gh extension install hushsecurity/gh-pr-revision
```

## Synopsis

```shell
gh pr-revision --help
```

```text
GitHub CLI extension for pull request revisions
v0.4.0
Usage: gh-pr-revision <command> [<args>]

Options:
  --help, -h             display this help and exit
  --version              display version and exit

Commands:
  create                 create revision
  diff                   compare revisions with 'git diff'
  difftool               compare revisions with 'git difftool'
  list                   list revisions
  show                   show revision
```

## Create a Revision

To create a new revision make sure your pull request is `OPEN` and is not `Draft`.

```shell
gh pr-revision create
```

This command assumes that the new revision is associated with the `HEAD` commit.
If the pull request tip commit is not the local `HEAD` specify a commitish
explicitly:

```shell
gh pr-revision create -c "<commitish>"
```

To attach a custom comment to the revision use `-e` to open your configured editor:

```shell
gh pr-revision create -e
```

## Aliases

The name of the extension is relatively long. Therefore, it may be helpful to create
aliases as follows:

```shell
gh alias set rvc 'pr-revision create'
gh alias set rvl 'pr-revision list'
gh alias set rvs 'pr-revision show'
```

Now a new revision can be created with a shorter command `gh rvc -e`.

## Limitations

1. `gh-pr-revision` stores some metadata in pull request comments. Subsequent revisions
   build on the metadata included in previous revisions. Therefore, deletion of a
   revision comment may affect correctness of following revisions.

1. The size of a revision comment is limited by the max size of a
   pull request comment allowed by GitHub. Therefore, if a pull request has a very
   long list of revisions, creation of a new one may fail due to GitHub limitations.
   This is because every new revision contains 2 additional comparison urls that make
   the metadata bigger.
