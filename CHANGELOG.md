<!-- markdownlint-configure-file { "MD024": { "siblings_only": true } } -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.7.2] - 2025-07-01

### Changed

- update dependencies to latest versions

## [0.7.1] - 2025-03-27

### Changed

- update dependencies to latest versions
- move to `golangci-lint v2`

## [0.7.0] - 2025-01-05

### Added

- Add `-n, --no-review` knob to `create` command.

  This is handy when a PR is already approved but needs to be rebased over a fresh
  `main` tip for merging. In this case we would like to add a revision because of the
  rebase, but there is no need to request reviews because in most cases the PR will be
  merged immediately after Build succeeds. If any changes are required after rebase
  another revision can be created including review requests. This revision has a clean
  diff because it doesn't include the rebase changes.

## [0.6.0] - 2024-12-16

### Added

- add `create -r <user-reviewer> -t <team-reviewer>` cli args to add reviewers at
  revision creation time. This is especially handy at creation of the first revision
  because one can omit the reviewers at PR creation and specify them only when
  first revision is published.

## [0.5.0] - 2024-12-09

### Added

- automatically refresh review requests when new revision is created.

  Note that it is impossible to automatically remove a user from reviewers list.
  Every new revision will post a review request to all users that were ever requested
  a review on current PR.

### Changed

- change Revision metadata to include a set of reviewers and team reviewers.
  Every revision's set is a union of all previous revisions' sets and the current set
  of outstanding reviews reported by GH Api.

## [0.4.0] - 2024-10-28

This release adds small usability features:

### Added

- add `-m, --message` parameter to `create` subcommand. This allows specifying a
  revision comment without opening an editor. If combined with the `-e` parameter,
  the message appears in the opened editor as pre-populated content.

### Changed

- `gh-pr-revision create` prints now the created revision number, and includes
  a url of the created PR link

## [0.3.0] - 2024-10-23

This release adds support for comparing revisions from CLI.

### Added

- add `diff FROM TO` subcommand, where `FROM,TO` are revisions to compare
- add `difftool FROM TO` subcommand, where `FROM,TO` are revisions to compare

## [0.2.0] - 2024-10-23

This release renames the tool from `gh-revision` to `gh-pr-revision`.

This was done because ideally we would like to have a `gh pr revision` command.
However, given that extensions cannot be put under builtin commands, we decided
to have the next best thing `gh pr-revision`.

Note that this release is not backward compatible. It changes the revision metadata
format. Therefore, old revisions will not be properly parsed.

### Changed

- rename the tool to `gh-pr-revision`.
- rename the repository to `hushsecurity/gh-pr-revision`
- cleanup the README

## [0.1.0] - 2024-10-20

### Added

Initial release.

[0.1.0]: https://github.com/hushsecurity/gh-pr-revision/releases/tag/v0.1.0
[0.2.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.1.0...v0.2.0
[0.3.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.2.0...v0.3.0
[0.4.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.3.0...v0.4.0
[0.5.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.4.0...v0.5.0
[0.6.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.5.0...v0.6.0
[0.7.0]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.6.0...v0.7.0
[0.7.1]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.7.0...v0.7.1
[0.7.2]: https://github.com/hushsecurity/gh-pr-revision/compare/v0.7.1...v0.7.2
