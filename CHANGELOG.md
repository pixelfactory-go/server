# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.5.0](https://github.com/pixelfactory-go/server/compare/v0.4.1...v0.5.0) (2026-01-21)


### Features

* add comprehensive fuzzing support ([#13](https://github.com/pixelfactory-go/server/issues/13)) ([1c3e95b](https://github.com/pixelfactory-go/server/commit/1c3e95befee867f6290f2c78682c7d372955f9ca))


### Bug Fixes

* run CodeQL SAST on all branches to improve security score ([1926e98](https://github.com/pixelfactory-go/server/commit/1926e98a621394d2dcd7c6cee57fffb9aa0af7d0))

## [0.4.1](https://github.com/pixelfactory-go/server/compare/v0.4.0...v0.4.1) (2026-01-21)


### Bug Fixes

* **ci:** add name to test job in CI workflow ([385239b](https://github.com/pixelfactory-go/server/commit/385239b729e3c60ead9e8d710ffffa4b9b5cfdfd))

## [0.4.0](https://github.com/pixelfactory-go/server/compare/v0.3.0...v0.4.0) (2026-01-19)


### Features

* modernize repository with documentation, CI/CD, and code quality improvements ([#4](https://github.com/pixelfactory-go/server/issues/4)) ([777d36a](https://github.com/pixelfactory-go/server/commit/777d36aeb9381ac1be7e410d5d5777d952e8c7e6))

## [Unreleased]

### Added
- Comprehensive README.md with usage examples and API documentation
- CONTRIBUTING.md with contribution guidelines and development workflow
- CHANGELOG.md for tracking project changes

### Changed
- Upgraded Go version to 1.25
- Enhanced CI/CD workflows with release automation
- Improved Makefile with better formatting checks

## [0.2.0] - 2024-01-16

### Added
- Default logger with name and port fields ([#3](https://github.com/pixelfactory-go/server/pull/3)) ([236c1d7](https://github.com/pixelfactory-go/server/commit/236c1d7))

### Changed
- Upgraded to Go 1.23 ([#2](https://github.com/pixelfactory-go/server/pull/2)) ([2594a9f](https://github.com/pixelfactory-go/server/commit/2594a9f))

### Added
- GoReleaser configuration for automated releases ([1eedb4d](https://github.com/pixelfactory-go/server/commit/1eedb4d))

## [0.1.0] - 2024-01-15

### Added
- Initial server implementation with graceful shutdown ([#1](https://github.com/pixelfactory-go/server/pull/1)) ([687898b](https://github.com/pixelfactory-go/server/commit/687898b))
- HTTP server with configurable timeouts
- TLS support for secure connections
- Signal handling for SIGTERM and SIGINT
- Functional options pattern for configuration
- Structured logging integration
- Example implementation

### Changed
- Updated configuration structure ([e5f624a](https://github.com/pixelfactory-go/server/commit/e5f624a))

[Unreleased]: https://github.com/pixelfactory-go/server/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/pixelfactory-go/server/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/pixelfactory-go/server/releases/tag/v0.1.0
