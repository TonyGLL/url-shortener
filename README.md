# Setting Up and Running Application

https://roadmap.sh/projects/url-shortening-service

## Prerequisites

Ensure you have the following tools installed on your system:

- **Golang** ([Download](https://go.dev/dl/))
- **Make** (Available by default on most Linux/macOS systems, installable on Windows via [Chocolatey](https://chocolatey.org/) or [Scoop](https://scoop.sh/))
- **Docker** ([Install Docker](https://docs.docker.com/get-docker/))
- **Docker Compose** ([Install Docker Compose](https://docs.docker.com/compose/install/))
- **Air** ([Install Air](https://github.com/cosmtrek/air)) - Live reload for Go applications

## Steps to Set Up the Application

### 1. Start PostgreSQL with Docker

Run the following command to start a PostgreSQL container:

```sh
make postgres
```

This command will:

- Pull the necessary PostgreSQL Docker image (if not already available)
- Start a PostgreSQL container using `docker-compose.yml`
- Expose the database on the configured port

### 2. Create the Database

After PostgreSQL is up, create the database by running:

```sh
make createdb
```

This will:

- Use the PostgreSQL instance to create the necessary database for the application

### 3. Apply Database Migrations

To apply all pending database migrations, run:

```sh
make migrateup
```

This will:

- Execute migration scripts to initialize or update the database schema

### 4. Make your own local.env in the root

```sh
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/url_shortener?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
VERSION=0.0.1
SECRET=secret
```

This will:

- Let you run the app correctly

### 5. Start the Development Server with Hot Reload

For continuous development with live reload, use:

```sh
make watch
```

or

```sh
CONFIG_FILE=local.env air
```

This will:

- Start the Golang application
- Watch for file changes and restart automatically

## Additional Commands

- **Revert Last Migration**

  ```sh
  make migratedown
  ```

  Rolls back the last applied migration.

- **Revert Last Migration**

  ```sh
  make dropdb
  ```

  Drop all the table from the postgreSQL database.

## Troubleshooting

### 1. Docker Container Fails to Start

- Ensure Docker is running.
- Check for port conflicts (default PostgreSQL runs on `5432`).
- Run `docker ps` to see if the container is already running.

### 2. Database Connection Issues

- Verify that PostgreSQL is running using `docker ps`.
- Check environment variables to ensure they match the database configuration.
- Try restarting the container: `make stop-postgres && make postgres`.

### 3. Migration Errors

- Ensure the `migrations` folder contains valid migration files.
- Check for syntax errors in migration scripts.

---

Now your Golang application should be up and running!