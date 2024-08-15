# Project Architecture Overview

## Layered Architecture

This project adopts a layered architecture pattern to improve maintainability, testability, and flexibility. This architecture divides the application into four clearly defined main layers:

1. **Presentation Layer** (Interfaces)
   - Responsible for interfaces with the external world (HTTP, gRPC, etc.)
   - Handles receiving user input and returning responses

2. **Application Layer** (Application)
   - Responsible for implementing use cases
   - Acts as a bridge between the domain layer and the presentation layer
   - Manages transactions and coordinates other layers

3. **Domain Layer** (Domain)
   - Defines business logic and rules
   - Contains entities and value objects
   - Independent of external concerns

4. **Infrastructure Layer** (Infrastructure)
   - Responsible for implementing technical details such as databases and external APIs
   - Provides technical support for other layers

Main advantages of this architecture:
- **Separation of Concerns**: Each layer has specific responsibilities, making code management easier.
- **Testability**: Each layer can be tested independently, making it easier to write unit tests.
- **Flexibility**: Clear dependencies between layers make future changes and extensions easier.
- **Scalability**: Each layer can be scaled out independently, making performance optimization easier.

## Project Directory Structure

The project directory structure reflects the layered architecture described above:

```
project-root/
├── main.go                 # Application entry point
├── internal/               # Project-specific non-public code
│   ├── domain/             # Domain layer: Business logic and entities
│   │   ├── schema/         # Schema-related domain logic
│   │   ├── field/          # Field-related domain logic
│   │   └── record/         # Record-related domain logic
│   ├── application/        # Application layer: Use cases and application logic
│   │   ├── schema/         # Schema-related application services
│   │   ├── dynamicapi/     # Dynamic API generation-related application services
│   │   └── validation/     # Validation-related services
│   ├── infrastructure/     # Infrastructure layer: Integration with external services
│   │   ├── database/       # Database-related implementations
│   │   ├── repository/     # Repository pattern implementations
│   │   └── cache/          # Cache-related implementations
│   └── interfaces/         # Interface layer: Interfaces with the external world
│       ├── http/           # HTTP-related implementations
│       │   ├── handler/    # HTTP request handlers
│       │   ├── middleware/ # HTTP middleware
│       │   └── router.go   # Routing configuration
│       └── grpc/           # gRPC-related implementations (for future expansion)
├── pkg/                    # General-purpose packages that can be used externally
│   ├── logger/             # Logging utilities
│   └── errors/             # Error handling utilities
├── config/                 # Application configuration-related
├── scripts/                # Development and operation scripts
├── test/                   # Test code
│   ├── unit/               # Unit tests
│   └── integration/        # Integration tests
├── docs/                   # Documentation
├── go.mod                  # Go module definition
└── go.sum                  # Go module dependency lock file
```

### Description of Key Directories

- `main.go`: The entry point of the application. Responsible for overall initialization and startup.
- `internal/`: Contains project-specific code and prevents external imports. The main implementation of the layered architecture is here.
  - `domain/`: Contains domain layer code. Defines business logic and entities.
  - `application/`: Contains application layer code. Defines use case implementations and application services.
  - `infrastructure/`: Contains infrastructure layer code. Includes database connections, repository implementations, and cache logic.
  - `interfaces/`: Contains presentation layer code. Includes HTTP handlers, middleware, and routing logic.
- `pkg/`: Contains reusable general-purpose packages. Can be imported by other projects.
- `config/`: Contains application configuration-related code.
- `scripts/`: Contains scripts necessary for development and deployment.
- `test/`: Contains unit tests and integration tests.
- `docs/`: Contains project-related documentation such as API documentation.

## Direction of Dependencies

One of the important principles of layered architecture is that dependencies should flow from outer layers to inner layers. Specifically:

1. The interface layer depends on the application layer.
2. The application layer depends on the domain layer.
3. The infrastructure layer may depend on all other layers, but not vice versa.

Maintaining this direction of dependencies allows us to keep the inner layers (especially the domain layer) independent of external concerns and stabilize the core of the system.

## Summary

Adopting this architecture and structure provides the following benefits:

1. Code Organization: Each component has a clearly defined place.
2. Maintainability: The impact of changes is limited, making code understanding and modification easier.
3. Testability: Each layer can be tested independently, making it easier to write unit tests.
4. Scalability: When adding new features or technologies, the impact on existing code can be minimized.
5. Team Development: Clear structure promotes communication and collaboration among team members.

This architecture can be flexibly adjusted according to project requirements and growth. It's important to review regularly and make improvements as needed.
