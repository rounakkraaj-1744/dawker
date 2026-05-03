# uldocker

A production-grade terminal-based Docker management tool.

## Features

- **Real-time Docker Interaction**: List, start, and stop containers.
- **Vim-inspired Navigation**: Keyboard-first interface.
- **Command Mode**: Enter `:` to run Docker commands directly.

## Getting Started

### Prerequisites

- Go (latest stable)
- Docker installed and running

### Installation

1. Clone the repository (or navigate to this folder).
2. Initialize and tidy dependencies:
   ```bash
   go mod tidy
   ```

### Running the App

```bash
go run cmd/main.go
```

## Usage

- `j/k` or `Up/Down`: Navigate container list.
- `:`: Enter command mode.
  - `:stop <container_name>`: Stop a container.
  - `:start <container_name>`: Start a container.
- `q` or `Ctrl+C`: Quit the application.
