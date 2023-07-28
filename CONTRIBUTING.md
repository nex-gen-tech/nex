## Contributing to `nex`

First off, thank you for considering contributing to `nex`! It's people like you that make `nex` such a great tool.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Issues](#issues)
- [Pull Requests](#pull-requests)
- [Testing](#testing)
- [Coding Style](#coding-style)
- [Commit Messages](#commit-messages)

## Code of Conduct

By participating in this project, you are expected to uphold our [Code of Conduct](CODE_OF_CONDUCT.md).

## Getting Started

### Setting Up Your Environment

1. Fork the `nex` repository to your own GitHub account.
2. Clone your fork to your local machine: `git clone https://github.com/YOUR_USERNAME/nex.git`
3. Navigate to the project directory: `cd nex`
4. Install the required dependencies: `go get -v ./...`
5. Build the project: `go build`

Now you're ready to make changes!

### Issues

- Before submitting a new issue, please check if a similar issue is already opened.
- If you find a bug, open an issue using the bug report template.
- For feature requests or enhancements, use the feature request template.
- Label the issue appropriately. If you're not sure which label to use, it's okay; the maintainers will help you out.
- Be as descriptive as possible. The more information you provide, the easier it is for maintainers and the community to understand and address the issue.

### Pull Requests

1. Create a new branch for your changes: `git checkout -b feature/my-new-feature`
2. Make your changes and commit them with a meaningful commit message.
3. Push your branch to your fork on GitHub: `git push origin feature/my-new-feature`
4. Open a pull request against the `nex` main repository.
5. Describe your changes in the pull request description and mention any related issues.
6. Ensure that your pull request passes all CI checks.
7. Wait for a review. Address any comments or feedback from the maintainers.
8. Once approved, your pull request will be merged into the main codebase.

### Testing

- Always write tests for your changes. Ensure that the project's test coverage remains high.
- Run tests locally using `go test ./...` before submitting a pull request.
- Check the CI status for your pull request to ensure all tests pass.

### Coding Style

- Follow the Go coding conventions.
- Use meaningful variable and function names.
- Comment your code when necessary, but prefer clear code over lots of comments.
- Keep functions and methods short and focused on a single task.

### Commit Messages

- Write clear and concise commit messages describing the changes you made.
- Use the present tense ("Add feature" not "Added feature").
- Reference issues or pull requests in your commits when relevant.

Thank you for your contributions to `nex`! Your efforts help make this project better for everyone.