# **Product Management System with Asynchronous Image Processing**

This project is a **Product Management System** with **asynchronous image processing** using **Go**, **PostgreSQL**, **Redis**, **RabbitMQ**, and **Docker**. It allows users to create, update, and retrieve products along with their associated images. The image processing happens asynchronously using a background service that listens to a message queue.

## **Architectural Overview**

### **1. Core Components**

The system is divided into several key components:

- **API Service**: 
  - The API service exposes the main application functionality (e.g., create and retrieve products) via RESTful endpoints. 
  - It is built using **Go** and structured to handle HTTP requests using the **Gorilla Mux** router.
  
- **Image Processing Service**: 
  - This service processes product images asynchronously by listening to a **RabbitMQ** queue.
  - Upon receiving a message with image details, it downloads the image, processes it (e.g., resizing or compression), and stores it in the database.
  
- **Database (PostgreSQL)**: 
  - **PostgreSQL** stores product information, including user IDs, product names, descriptions, and image data.
  
- **Redis**: 
  - **Redis** is used as a cache to store frequently accessed product information, reducing the load on the database.
  
- **RabbitMQ**: 
  - **RabbitMQ** is used to handle the asynchronous image processing. The image processing service consumes messages from a queue to process product images in the background.

### **2. Folder Structure**

The project follows a modular folder structure to separate concerns clearly:

```
product-management-system/
├── api/
│   ├── handlers/                # API handler logic (product-related operations)
│   ├── router.go                # Routes API endpoints to the appropriate handler
│   └── main.go                  # API entry point
├── async/
│   ├── image_processing/        # Image processing service logic
│   │   ├── processor.go         # Image processing logic (downloads, resizes, uploads images)
│   │   └── queue.go             # RabbitMQ queue consumption for image processing
│   └── main.go                  # Entry point for the async service
├── caching/
│   └── redis.go                 # Redis caching logic
├── database/
│   ├── db.go                    # Database connection and schema setup
│   └── migrations/              # Database schema migrations
├── models/
│   ├── product.go               # Product model (schema)
│   ├── user.go                  # User model (schema)
│   └── image.go                 # Image model (processed images)
├── logging/
│   └── log.go                   # Logging setup (structured logging)
├── tests/
│   ├── api_tests/               # Unit and integration tests for API endpoints
│   ├── async_tests/             # Unit tests for async services
│   └── caching_tests/           # Tests for Redis caching functionality
├── config/
│   └── config.yaml              # Configuration for DB, Redis, RabbitMQ, etc.
├── Dockerfile                   # Dockerfile to build the API service
├── docker-compose.yml           # Docker Compose configuration to run services
├── README.md                    # Project documentation
└── .env                          # Environment variables (sensitive data like DB passwords)
```

### **3. Key Technologies Used**
- **Go (Golang)**: A statically typed, compiled language known for its performance, ease of use in building APIs, and strong concurrency support.
- **PostgreSQL**: A powerful, open-source relational database system to store product and user data.
- **Redis**: An in-memory key-value store used for caching product data.
- **RabbitMQ**: A message broker that handles the queueing of tasks for asynchronous image processing.
- **Docker**: Used to containerize the application, making it portable and easy to run across different environments.
- **Docker Compose**: Orchestrates multi-container applications, helping to manage services like PostgreSQL, Redis, RabbitMQ, and the API.

## **Setup Instructions**

### **1. Clone the Repository**

Clone the project to your local machine:

```bash
git clone https://github.com/your-username/product-management-system.git
cd product-management-system
```

### **2. Environment Setup**

Ensure that you have the following installed:
- **Docker** and **Docker Compose** for containerization and orchestration.
- **Go 1.19** or later for building the application (if running without Docker).
- **PostgreSQL**, **Redis**, and **RabbitMQ** if you're running them locally (otherwise, Docker Compose will handle them for you).

### **3. Configuration File (`config/config.yaml`)**

Create or modify the `config/config.yaml` file to include your database, Redis, and RabbitMQ configurations:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "example"
  dbname: "product_db"
  sslmode: "disable"

redis:
  host: "localhost"
  port: 6379

rabbitmq:
  host: "localhost"
  port: 5672
  username: "guest"
  password: "guest"

log:
  level: "info"  # Can be "debug", "info", "warn", "error"
```

### **4. Build and Run with Docker Compose**

Docker Compose is used to build and run the services (API, PostgreSQL, Redis, RabbitMQ) together.

To build and run the application:

1. Make sure your terminal is in the project root directory (where the `docker-compose.yml` file is located).
2. Run the following command to start all the services:

```bash
docker-compose up --build
```

This will:
- Build the Docker images for the **API service** and **Image Processing service**.
- Start **PostgreSQL**, **Redis**, **RabbitMQ**, and the **API service**.
- The **API service** will be available at `http://localhost:8080`.
- The **Image Processing service** will be available (if required) on `http://localhost:8081`.

### **5. Verify Services**

- **API**: Test the API by sending requests to `http://localhost:8080` using **Postman** or **cURL**.
- **Database**: PostgreSQL will be available at `localhost:5432`. You can access it using a client like **pgAdmin** or **psql**.
- **Redis**: Redis will be available at `localhost:6379`.
- **RabbitMQ**: RabbitMQ management interface will be available at `http://localhost:15672` (default credentials are `guest/guest`).

### **6. Running without Docker (Optional)**

If you prefer to run the application locally without Docker, you'll need to manually install the required services (PostgreSQL, Redis, RabbitMQ). After that, you can:

1. Run `go mod tidy` to install the required dependencies.
2. Run `go run api/main.go` to start the API service.

## **Assumptions**
- The project is designed to be run inside Docker for easy service management and scalability.
- The configuration in `config.yaml` is assumed to be correct. Adjust it based on your environment if running outside Docker.
- The **image processing** service communicates with RabbitMQ for background tasks (like downloading and processing images). RabbitMQ and Redis are expected to be available for the services to interact with.
- The **PostgreSQL** database is used for storing product data, and Redis is used for caching frequent queries to reduce load on the database.
- The **API** service communicates with the database to handle product CRUD operations and responds with JSON data.

## **Testing**

### **1. Unit Tests**

The project includes unit tests for various components:
- **API tests**: Located in `tests/api_tests/`.
- **Async tests**: Located in `tests/async_tests/`.
- **Redis caching tests**: Located in `tests/caching_tests/`.

To run tests:

```bash
go test ./tests/
```

### **2. Integration Tests**

You can also perform integration tests to verify end-to-end functionality, especially for background tasks and cache hits.

---

## **Troubleshooting**

- **Docker Build Issues**: If you encounter build issues related to Docker, ensure that the `docker-compose.yml` file and all paths are correct.
- **Database Connection**: Make sure the database credentials in `config.yaml` match those configured in the `docker-compose.yml` file for PostgreSQL.
- **Service Not Starting**: Check the logs for any error messages. Run `docker-compose logs` to view logs for all services.

---

## **Future Enhancements**

- **User Authentication**: Add user authentication and authorization to the API.
- **Image Optimization**: Enhance the image processing service to include more advanced image optimization techniques.
- **Microservices**: Scale out the architecture by splitting services further into distinct microservices.

---

### **Conclusion**

This project is designed to demonstrate a scalable product management system with asynchronous image processing. By leveraging Docker, RabbitMQ, Redis, and PostgreSQL, we have created a robust architecture that ensures efficient product management and background task handling.

Let me know if you need any further information or if you'd like assistance with additional setup!