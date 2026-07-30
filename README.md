# CP Web Application Starter Template

This repository contains a full-stack starter with a Go Fiber backend and a Next.js frontend. The backend is structured around dependency-injected services and router-scoped endpoint packages.

## Repository structure

```text
.
├── docker-compose.yml
├── README.md
├── src/
│   ├── backend/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── main_test.go
│   │   └── internal/
│   │       ├── endpoints/
│   │       │   ├── healthz/
│   │       │   │   ├── endpoint.go
│   │       │   │   └── router.go
│   │       │   └── foobar/
│   │       │       ├── endpoint.go
│   │       │       └── router.go
│   │       ├── router/
│   │       │   ├── router.go
│   │       │   └── router_test.go
│   │       └── service/
│   │           ├── service.go
│   │           ├── service_mock.go
│   │           └── service_test.go
│   └── frontend/
│       ├── app/
│       ├── components/
│       ├── package.json
│       └── ...
└── .github/workflows/
    ├── ci.yml
    ├── backend-quality-gate.yml
    ├── backend-tests.yml
    ├── frontend-quality-gate.yml
    └── frontend-tests.yml
```

## Backend architecture

- The endpoint handlers live in `src/backend/internal/endpoints/<router-name>/`.
- Each router folder contains its own `endpoint.go` and `router.go` files.
- The top-level router package wires those endpoint packages into the Fiber app.
- The service layer is dependency-injected and composed in `src/backend/internal/service/service.go`.

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

Available backend routes:

- `GET /healthz`
- `GET /foo/bar`

You can also run the test suite with:

```bash
go test ./...
```

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

GitHub Actions workflow definitions live in `.github/workflows/`. The main entrypoint is `.github/workflows/ci.yml`, and the repository also includes backend and frontend quality gate workflows for future CI expansion.

### Suggested next checks

- Backend checks:
  - `go vet ./...`
  - `go test ./...`
  - `golangci-lint`
- Frontend checks:
  - ESLint
  - unit tests
  - coverage thresholds
  - npm audit

If you want stricter enforcement, enable the relevant workflow jobs in `.github/workflows/ci.yml` and protect the main branch so pull requests cannot be merged until the checks pass.