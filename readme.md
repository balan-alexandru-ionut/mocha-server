# Mocha Server

Mocha is an application for managing personal tasks and projects which can be
self-hosted.

This repo hosts the backend of the application.

## Prerequisites:

Have the following dependencies installed:
- Go - `1.22.x` or later
- Docker - `25.0.3` or later

## 1. Application setup:

### 1.1. Database setup

Mocha uses MongoDB as a database and provides a `docker-compose.yaml` file for quickly spinning up a database container.

To start, edit the variables for username and password in the `database-setup.sh` script. This will export two environment variables which will be used by the Go application and the MongoDB cotainer.

These variables will be available until the next restart of your machine. If you want to persist them between restarts you can add the contents of the script to you shell rc file.

To run the script execute:

```bash
chmod +x ./database-setup.sh

source database-setup.sh
```

### 1.2. Start database container

```bash
docker compose up -d
```

### 1.3. Build the application

```bash
go build
```

## 2. Start the application:

```bash
./mocha-server
```

## 3. TODO: hosting with caddy

## 4. TODO: One command docker setup of application, proxy and database