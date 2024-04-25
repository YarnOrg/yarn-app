# Core Yarn App

## Overview

Yarn is a social finance app combines the functionality of a social media platform with financial transactions, allowing users to chat, make payments, and participate in a marketplace environment. This project is structured as a modular monolith, using Golang for the backend services.

### Local Prerequisites

What things you need to install the software and how to install them:

- [Go](https://golang.org/dl/) (version 1.21 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)

### Local setup via app or docker

```bash
   git clone https://github.com/YarnOrg/yarn-app.git
   cd yarn-app
   go mod tidy
   cp .env.example .env
   go run cmd/api/main.go
```

```
docker build -t yarn-app .
docker run -p 8080:8080 yarn-app

or with env vars

docker run -p 8080:8080 yarn-app -e ....
```

## Directory Structure

```
.
├── README.md
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── internal
│   ├── chat
│   ├── marketplace
│   ├── transaction
│   ├── user
│   └── wallet
├── pkg
│   ├── config
│   ├── db
│   └── util
└── scripts
```
