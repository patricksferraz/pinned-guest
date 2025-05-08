# Pinned Guest

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-guest)](https://goreportcard.com/report/github.com/patricksferraz/pinned-guest)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-guest?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-guest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, scalable microservice built with Go, featuring a clean architecture design and robust infrastructure setup.

## 🚀 Features

- **Clean Architecture**: Well-structured codebase following clean architecture principles
- **REST API**: Fast and efficient REST endpoints using Fiber framework
- **Database Support**:
  - PostgreSQL integration with GORM
  - SQLite support for development
  - MongoDB integration
- **Message Queue**: Kafka integration for event-driven architecture
- **Containerization**: Docker support with docker-compose for easy deployment
- **Kubernetes Ready**: Includes k8s configuration files for production deployment
- **Hot Reload**: Development environment with Air for hot reloading
- **API Documentation**: Swagger/OpenAPI documentation
- **Environment Configuration**: Flexible configuration management with environment variables

## 🛠️ Tech Stack

- **Language**: Go 1.18+
- **Web Framework**: Fiber v2
- **Database ORM**: GORM
- **Message Broker**: Apache Kafka
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Documentation**: Swagger/OpenAPI
- **Development**: Air (hot reload)

## 📋 Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose
- Make (for using Makefile commands)
- PostgreSQL (for production)
- Kafka (for message queue functionality)

## 🚀 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/patricksferraz/pinned-guest.git
   cd pinned-guest
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Run with Docker**
   ```bash
   make docker-up
   ```

4. **Run locally**
   ```bash
   make run
   ```

## 🏗️ Project Structure

```
.
├── app/            # Application layer
│   └── rest/       # REST API handlers
├── cmd/            # Command line interface
├── domain/         # Domain models and interfaces
├── infra/          # Infrastructure implementations
├── k8s/            # Kubernetes configurations
└── utils/          # Utility functions
```

## 🔧 Development

- **Run tests**
  ```bash
  make test
  ```

- **Run linter**
  ```bash
  make lint
  ```

- **Generate API documentation**
  ```bash
  make swagger
  ```

## 📚 API Documentation

Once the application is running, you can access the Swagger documentation at:
```
http://localhost:{REST_PORT}/swagger/
```

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/patricksferraz/pinned-guest/issues).

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Sferraz** - *Initial work* - [patricksferraz](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/go-gorm/gorm)
- [Air](https://github.com/cosmtrek/air)
- And all other open-source projects that made this possible!
