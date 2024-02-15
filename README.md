# renderhook

Industry-leading speed and reliable API-driven image generation, with zero cold start and minimal latency.

## Getting Started

To get started with this project, you can use the provided Makefile commands to install dependencies, run the application in development mode, build the application for production, and run tests.

### Prerequisites

- Go installed for the backend.
- Node.js and Bun package manager for the frontend.

### All in One

To install all dependencies, start the application in development mode, build the application for production, and run tests, run:

```bash
make
```

This uses `concurrently` to run both parts of the application simultaneously.

### Installation

Run the following command to install all dependencies for both backend and frontend:

```bash
make install
```

### Development

To start both the backend and frontend in development mode, run:

```bash
make dev
```

This uses `concurrently` to run both parts of the application simultaneously.

### Building

To build both the backend and frontend for production, run:

```bash
make build
```

### Testing

To run tests for both backend and frontend, run:

```bash
make test
```

## Makefile Targets

The Makefile includes several targets for managing the application:

- `all`: Installs dependencies and starts both backend and frontend in development mode.
- `install`: Installs all necessary dependencies for both backend and frontend.
- `dev`: Starts both the backend and frontend in development mode.
- `build`: Builds both the backend and frontend for production.
- `test`: Runs tests for both backend and frontend.
- `clean`: Cleans up build artifacts for both backend and frontend.

### Backend Tasks

- `be-install`: Installs dependencies for the Go backend.
- `be-dev`: Starts the Go backend in development mode.
- `be-build`: Builds the Go backend for production.
- `be-test`: Runs tests for the Go backend.

### Frontend Tasks

- `fe-install`: Installs dependencies for the NextJS frontend.
- `fe-dev`: Starts the NextJS frontend in development mode.
- `fe-build`: Builds the NextJS frontend for production.
- `fe-test`: Runs tests for the NextJS frontend.

## Clean Task

The `clean` task can be used to remove build artifacts and clean up the project directory.
