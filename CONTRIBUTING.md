# Contributing to Go API Boilerplate

Thank you for considering contributing to this project! This document provides guidelines for contributing.

## Code of Conduct

Please be respectful and constructive in your interactions with other contributors.

## How to Contribute

### Reporting Bugs

- Check if the bug has already been reported in Issues
- If not, create a new issue with a clear description
- Include steps to reproduce the bug
- Provide system information and error messages

### Suggesting Enhancements

- Open an issue with the enhancement proposal
- Explain why this enhancement would be useful
- Provide examples of how it would work

### Pull Requests

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes following the code style
4. Add or update tests as needed
5. Ensure all tests pass
6. Update documentation if necessary
7. Commit with clear, descriptive messages
8. Push to your fork and submit a pull request

## Development Setup

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-api-boilerplate.git
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file from `.env.example`

4. Run tests:
   ```bash
   make test
   ```

## Code Style

- Follow Go conventions and idioms
- Run `go fmt` before committing
- Run `golangci-lint` to check for issues
- Write clear, descriptive variable and function names
- Add comments for complex logic
- Keep functions small and focused

## Testing

- Write unit tests for new features
- Ensure existing tests pass
- Aim for good test coverage
- Use table-driven tests where appropriate

## Commit Messages

- Use clear, descriptive commit messages
- Start with a verb in present tense (Add, Fix, Update, etc.)
- Keep the first line under 72 characters
- Add detailed description if necessary

Example:
```
Add user authentication feature

- Implement JWT token generation
- Add middleware for protected routes
- Update documentation
```

## Architecture Guidelines

Follow the Clean Architecture principles:

- **Handler Layer**: HTTP handling only, no business logic
- **Service Layer**: Business logic, validation, orchestration
- **Repository Layer**: Database operations only

### Adding New Features

1. Define models in `internal/model/`
2. Create repository interface and implementation
3. Implement service layer with business logic
4. Create handlers for HTTP endpoints
5. Register routes in `cmd/api/main.go`
6. Add migrations if database changes are needed
7. Write tests for all layers
8. Update documentation

## Documentation

- Update README.md for user-facing changes
- Add code comments for complex logic
- Update API documentation for new endpoints
- Include examples where helpful

## Questions?

Feel free to open an issue for questions or clarifications.

Thank you for contributing!
