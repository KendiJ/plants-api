# Home Plants APIs
* a Golang project for practice, fun and all cool stuff I guess 


A robust backend service built with **Go (Golang)** and **MySQL**, designed to manage a household plant inventory across different rooms. This project demonstrates clean architecture, RESTful API design, and containerized database management.

## Tech Stack & Tools
* **Language:** Go (1.25+)
* **Database:** MySQL 8.0 (Running in Docker)
* **Architecture:** Standard Go Project Layout (cmd/internal pattern)
* **Routing:** Go Standard Library (`net/http` with Go 1.22+ routing)
* **Containerization:** Docker Desktop

## Project Structure
The architecture:
```text
homeplants-api/
├── cmd/api/main.go          # Entry point (wires everything together)
├── internal/
│   ├── database/            # DB connection & Migration logic
│   ├── handlers/            # HTTP Controllers (Business logic)
│   └── models/              # Structs & JSON tagging
├── schema.sql               # Database schema & Seed data
└── .env                     # Secrets (Gitignored)