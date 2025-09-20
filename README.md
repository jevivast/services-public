# Clean Architecture Go API

This is a Go API project following Clean Architecture principles.

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── application/         # Application business rules
│   ├── domain/              # Enterprise business rules
│   │   ├── models/          # Domain models
│   │   ├── ports/           # Interfaces (ports)
│   │   └── services/        # Domain services
│   └── infrastructure/      # Frameworks & drivers
│       ├── config/          # Configuration
│       ├── handlers/        # HTTP handlers
│       └── repositories/    # Data access layer
├── pkg/                     # Reusable packages
│   ├── middleware/          # HTTP middleware
│   └── utils/               # Utility functions
├── api/                     # API definitions
│   ├── v1/                  # API version 1
│   └── health/              # Health check endpoints
├── configs/                 # Configuration files
├── scripts/                 # Build and deployment scripts
└── docs/                    # Documentation
```

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

- `GET /api/v1/health` - Health check endpoint

## Configuration

Configuration is managed through YAML files in the `configs` directory. The application looks for a `config.yaml` file by default.
