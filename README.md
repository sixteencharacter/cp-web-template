# CP Web Application Starter Template

This repository contains a simple full-stack starter with a Go Fiber backend and a Next.js frontend.

## Project structure

- src/backend: Go API server
- src/frontend: Next.js app
- .github/workflows: CI workflow definitions

## Prerequisites

Before running the project locally, install:

- Go 1.26.5 or newer
- Node.js 20 LTS
- npm

## Backend setup

1. Open a terminal and move to the backend folder:
   ```bash
   cd src/backend
   ```
2. Download Go dependencies:
   ```bash
   go mod download
   ```
3. Start the backend:
   ```bash
   go run .
   ```
4. Open http://localhost:3000 to verify the server is responding.

> Note: the backend currently listens on port 3000. If you want to run the frontend at the same time, start the frontend on another port as shown below.

## Frontend setup

1. Open a second terminal and move to the frontend folder:
   ```bash
   cd src/frontend
   ```
2. Install frontend dependencies:
   ```bash
   npm ci
   ```
3. Start the frontend development server:
   ```bash
   npm run dev -- --port 3001
   ```
4. Open http://localhost:3001 in your browser.

## Run both apps locally

Use two terminals:

```bash
# Terminal 1
cd src/backend
go run .
```

```bash
# Terminal 2
cd src/frontend
npm ci
npm run dev -- --port 3001
```

## CI overview

GitHub Actions is configured in .github/workflows/ci.yml. The current pipeline runs on every push and pull request.

### Current workflow

The active workflow calls .github/workflows/pre.yml, which:

- formats Go's backend code with gofmt
- formats frontend files with Prettier
- commits the formatting changes back to the branch if the code is malformed
- fails the workflow if formatting changes were needed

In practice, CI is currently focused on code style and formatting rather than full quality checks.

### What is already prepared but not enabled

The repository also contains these workflow files:

- .github/workflows/frontend-quality-gate.yml
- .github/workflows/frontend-tests.yml

These workflows are ready to add linting, security scans, test execution, and coverage checks, but they are currently commented out in ci.yml and can be edited later.

## How to expand CI with more checks

1. Open .github/workflows/ci.yml.
2. Uncomment the jobs you want to enable.
3. Make sure the required tools or scripts exist in the frontend package. For example, the current test workflow expects a test:coverage script, but it is not defined in the frontend package yet.

Example:

```yaml
jobs:
  pre-check:
    uses: ./.github/workflows/pre.yml

  frontend-quality-gates:
    needs:
      - pre-check
    uses: ./.github/workflows/frontend-quality-gate.yml

  tests:
    needs:
      - pre-check
    uses: ./.github/workflows/frontend-tests.yml
```

### Good next checks to add

- Backend checks:
  - go vet ./...
  - go test ./...
  - golangci-lint
- Frontend checks:
  - ESLint
  - unit tests
  - coverage thresholds
  - npm audit
  - Semgrep or Gitleaks

### Practical suggestion

Start by enabling the existing frontend quality gate and test workflows, then add backend checks once the project has a few real Go tests.

If you want stricter enforcement, also enable branch protection rules in GitHub so pull requests cannot be merged until the required checks pass.