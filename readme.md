<p align="center">
<img src="https://openupthecloud.com/wp-content/uploads/2020/01/Golang.png" alt="Go Logo" width="300">
</p>
# API-Golang Project

This project is a RESTful API built with Go (Golang), following best practices and clean architecture principles. It features a structured project layout, dependency injection, unit testing, and more.

## Project Structure

The project is organized into various packages:

- `bootstrap`: Bootstrapping and dependency injection setup.
- `config`: Configuration management.
- `controller`: HTTP handlers that process requests and return responses.
- `database`: Database connection setup and initialization.
- `dto`: Data Transfer Objects for input validation.
- `errors`: Custom error definitions and handling.
- `migrations`: Database migrations for evolving the database schema.
- `models`: Data models representing the business entities.
- `repositories`: Data access layer providing an abstraction over the database interactions.
- `services`: Business logic layer.
- `services/mocks`: Mock services for unit testing.

## Getting Started

### Prerequisites

- Go 1.15 or higher
- PostgreSQL
- A proper Go workspace setup

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/api-golang.git
```

2. Navigate to the project directory:
```sh
cd api-golang
```

3. Install the dependencies:
```sh
go mod tidy
```

4. Set up environment variables by creating a `.env` file in the project root with the following content:
```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_NAME=your_db_name
DB_PORT=5432
DB_SSLMODE=disable
```

5. Run the database migrations:
```sh
make migrateup
```

### Using Air for Hot Reloading

`Air` is a tool for live reloading your Go applications, allowing for automatic rebuilding and restarting whenever file changes are detected, significantly improving your development workflow.

1. Install air by running:
```sh
go install github.com/cosmtrek/air@latest
```

2. Create a configuration file named `.air.toml` in your project root, or use the default configuration by running `air` init.

3. Start the development server with air: 
```sh
air
```

### Running the Server

Execute the tests using the following command:
```sh
go test ./...
```