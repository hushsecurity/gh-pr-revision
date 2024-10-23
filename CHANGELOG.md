<!-- markdownlint-configure-file { "MD024": { "siblings_only": true } } -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
