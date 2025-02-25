# GoSimpleBank

## Overview
Go Simple Bank is a backend service for managing bank accounts, tracking balances, and enabling secure money transfers. The backend is built with **Golang**, **PostgreSQL**, **Gin**, **gRPC**, and **Kubernetes**, ensuring performance, security, and scalability.

### Key Development Areas

### 1. Database & Backend Management
- PostgreSQL schema design, migrations, transactions, and isolation levels  
- Handling deadlocks, using **pgx** for optimized database access  
- Implementing authentication with **PASETO/JWT** and enforcing RBAC-based access control  
- Background workers and async task processing with **Redis + Asynq**  
- Secure session management and CORS policy enforcement  

### 2. API Development & Communication
- Developing RESTful APIs with **Gin**, implementing middleware for access control  
- Creating gRPC APIs with **Protobuf**, extracting metadata, and integrating Swagger documentation  
- Mocking databases for unit testing and error handling  

### 3. Deployment & Scalability
- Containerizing with **Docker**, managing secrets with **AWS Secrets Manager**  
- Deploying to **AWS EKS** using **Kubernetes**, setting up **AWS RDS**  
- Configuring **ingress, TLS encryption,** and **auto-scaling**  
- Ensuring graceful server shutdown and optimizing AWS costs  

### 4. Security, CI/CD & Observability
- Implementing CI/CD pipelines with **GitHub Actions**  
- Logging, monitoring, and error handling  
- Sending verification emails via Gmail

## Features
- **Account Management:** Create and manage bank accounts with owner name, balance, and currency.
- **Transaction History:** Record every balance change to ensure accurate tracking.
- **Money Transfers:** Securely transfer funds between accounts using transactions to ensure consistency.

## Technology Stack
- **Language:** Golang  
- **Database:** PostgreSQL
- **Web Framework:** Gin  
- **Authentication:** PASETO, JWT, Bcrypt  
- **Deployment:** Docker, Kubernetes, AWS (EKS, RDS, ECR, Route53)  
- **Async Processing:** Redis, Asynq  
- **CI/CD:** GitHub Actions  

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
  
  Using [scoop](https://scoop.sh/)

  ```bash
  scoop install migrate
  ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

  ```bash
  scoop install sqlc
  ```

- [Gomock](https://github.com/golang/mock)

  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup infrastructure

- Create the bank-network

  ```bash
  make network
  ```

- Start postgres container:

  ```bash
  make postgres
  ```

- Create simple_bank database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup
  ```

- Run db migration up 1 version:

  ```bash
  make migrateup1
  ```

- Run db migration down all versions:

  ```bash
  make migratedown
  ```

- Run db migration down 1 version:

  ```bash
  make migratedown1
  ```

### How to generate code

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```

- Generate DB mock with gomock:

  ```bash
  make mock
  ```

- Create a new db migration:

  ```bash
  make new_migration name=<migration_name>
  ```

### How to run

- Run server:

  ```bash
  make server
  ```

- Run test:

  ```bash
  make test
  ```
