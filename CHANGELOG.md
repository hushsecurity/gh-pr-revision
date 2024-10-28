<!-- markdownlint-configure-file { "MD024": { "siblings_only": true } } -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
