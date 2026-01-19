# Contributing to Server Package

Thank you for your interest in contributing to the Server package! We welcome contributions from the community and are grateful for your support.

## Code of Conduct

We are committed to providing a welcoming and inclusive environment for all contributors. Please maintain a respectful and inclusive environment for all contributors, regardless of their background or experience level.

## How to Contribute

### Reporting Issues

Before submitting an issue, please:

1. **Search existing issues** to avoid duplicates
2. **Use the GitHub issue tracker** for bug reports and feature requests
3. **Provide detailed information** including:
   - Go version and operating system
   - Steps to reproduce the issue
   - Expected vs actual behavior
   - Relevant code snippets or error messages
   - Server configuration details

### Pull Request Process

1. **Fork the repository** and create your branch from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Write or update tests** for your changes:
   ```bash
   make test
   ```

4. **Run linting** to ensure code quality:
   ```bash
   make lint
   ```

5. **Format your code** using gofmt:
   ```bash
   make fmt
   ```

6. **Commit your changes** using conventional commits (see below)

7. **Push to your fork** and submit a pull request to the `main` branch

8. **Respond to review feedback** and make any necessary updates

## Development Setup

### Prerequisites

- **Go 1.25 or higher**
- **golangci-lint** for code linting
  ```bash
  # Install golangci-lint
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

### Local Development

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/server.git
   cd server
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests:
   ```bash
   make test
   ```

4. Run linter:
   ```bash
   make lint
   ```

## Coding Standards

### Go Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` for code formatting (enforced by `make fmt`)
- Write idiomatic Go code that is clear and maintainable
- Keep functions focused and single-purpose
- Use meaningful variable and function names

### Documentation

- Document all exported functions, types, and constants
- Use godoc-compatible comments
- Include code examples for complex functionality
- Update README.md if adding new features

### Testing

- Write unit tests for all new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Include both positive and negative test cases
- Test edge cases and error conditions

Example table-driven test:

```go
func TestServer(t *testing.T) {
    tests := []struct {
        name    string
        options []Option
        wantErr bool
    }{
        {
            name:    "default server",
            options: nil,
            wantErr: false,
        },
        {
            name: "custom port",
            options: []Option{
                WithPort("3000"),
            },
            wantErr: false,
        },
        // Add more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            srv, err := New(tt.options...)
            if (err != nil) != tt.wantErr {
                t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
            }
            if !tt.wantErr && srv == nil {
                t.Error("New() returned nil server")
            }
        })
    }
}
```

## Commit Message Convention

We follow [Conventional Commits](https://www.conventionalcommits.org/) for commit messages. This helps with automated changelog generation.

### Format

```
<type>: <description>

[optional body]

[optional footer]
```

### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **refactor**: Code refactoring without feature changes
- **test**: Adding or updating tests
- **chore**: Maintenance tasks, dependency updates
- **perf**: Performance improvements
- **ci**: CI/CD changes

### Examples

```
feat: add support for custom TLS configuration

Add WithTLSConfig option to enable HTTPS support with custom
TLS settings. Includes certificate loading and validation.
```

```
fix: prevent panic on nil logger

Check for nil logger before logging shutdown events to avoid
runtime panic during graceful shutdown.
```

```
docs: update README with TLS examples

Add comprehensive TLS/HTTPS server examples showing certificate
loading and configuration options.
```

## Pull Request Guidelines

### Before Submitting

- [ ] Tests pass locally (`make test`)
- [ ] Linting passes (`make lint`)
- [ ] Code is formatted (`make fmt`)
- [ ] Documentation is updated
- [ ] Commit messages follow conventional commits
- [ ] Branch is up to date with `main`

### PR Description

Your pull request should include:

1. **Clear title** following conventional commit format
2. **Description** of what the PR does and why
3. **Related issues** referenced (e.g., "Fixes #123")
4. **Testing performed** and results
5. **Breaking changes** clearly noted (if any)

### Review Process

1. **Automated checks** must pass (CI/CD)
2. **Code review** by maintainers
3. **Feedback incorporation** as needed
4. **Final approval** before merging
5. **Squash and merge** to maintain clean history

## Questions or Problems?

- **Documentation**: Check the [README.md](README.md) first
- **Issues**: Search existing issues or create a new one
- **Discussions**: Use GitHub Discussions for questions and ideas

## License

By contributing, you agree that your contributions will be licensed under the MIT License. All contributions will fall under the project's existing MIT License.

---

Thank you for contributing to the Server package! 🎉
