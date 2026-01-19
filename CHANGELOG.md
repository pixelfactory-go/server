# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
