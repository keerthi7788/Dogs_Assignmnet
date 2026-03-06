# Dogs_Assignmnet API's

This project implements a REST API for managing a list of dogs. The service allows users to create, read, update, and delete dog records through HTTP endpoints.

The application is built using Go with the Chi router and uses PostgreSQL for persistent data storage. The service is containerized with Docker and deployed publicly using Render so it can be accessed over the internet.

The goal of this project is to expose a list of dogs through an HTTP API and allow modifications to the list while ensuring that all interactions are stored permanently in the database.

------------------------------------------------------------
Clone the repository

https://dogs-assignmnet-1.onrender.com
Live Application

Public URL

https://dogs-assignmnet-1.onrender.com

Health Check Endpoint

https://dogs-assignmnet-1.onrender.com/api/v1/health

Example Response

{
  "message": "!!! We are RunninGoo !!!"
}

------------------------------------------------------------

Technology Stack

Language: Go (Golang)
Router: Chi Router
Database: PostgreSQL
Containerization: Docker
Deployment: Render

------------------------------------------------------------

Architecture

The project follows a three layer architecture.

Hndler Layer  
Handles routing, request parsing, and response formatting.

Service Layer  
Contains the business logic and validations.

Repository Layer  
Responsible for database operations and interacting with PostgreSQL.

Flow

Client → HTTP Handlers → Service Layer → Repository Layer → PostgreSQL

This separation helps keep the code modular and easier to maintain.

------------------------------------------------------------

Project Structure

Dogs_Asiignment

cmd/dogs-service
  main.go
  Logger
  models

http
  handlers
  response
  server.go

service
  business logic

repository
  postgres implementation

logger
  logging utilities
dogs.json
Dockerfile
docker-compose.yml
README.md

------------------------------------------------------------

Running the Application Locally

Clone the repository

https://github.com/keerthi7788/Dogs_Assignmnet.git

cd dogs-service

Build the Docker image

docker build -t dogs-service .

Run the container

docker run -p 8080:8080 dogs-service

The application will start on

http://localhost:8082

------------------------------------------------------------

API Endpoints

Base Path

/api/v1

------------------------------------------------------------

Create Dog

POST /api/v1/dogs

Request Body

{
  "breed": "hound",
  "sub_breed": "afghan"
}

Example curl

curl -X POST https://dogs-assignmnet-1.onrender.com/api/v1/dogs \
-H "Content-Type: application/json" \
-d '{
"breed":"hound",
"sub_breed":"afghan"
}'

Example Response

{
  "id": 1,
  "breed": "hound",
  "sub_breed": "afghan"
}

------------------------------------------------------------

Get All Dogs

GET /api/v1/dogs

Example curl

curl https://dogs-assignmnet-1.onrender.com/api/v1/dogs

Example Response

[
  {
    "id": 1,
    "breed": "hound",
    "sub_breed": "afghan"
  },
  {
    "id": 2,
    "breed": "bulldog",
    "sub_breed": "english"
  }
]

------------------------------------------------------------

Get Dog By ID

GET /api/v1/dogs/{id}

Example

GET /api/v1/dogs/1

Example curl

curl https://dogs-assignmnet-1.onrender.com/api/v1/dogs/1

Example Response

{
  "id": 1,
  "breed": "hound",
  "sub_breed": "afghan"
}

------------------------------------------------------------

Update Dog

PUT /api/v1/dogs/{id}

Request Body

{
  "breed": "retriever",
  "sub_breed": "golden"
}

Example curl

curl -X PUT https://dogs-assignmnet-1.onrender.com/api/v1/dogs/1 \
-H "Content-Type: application/json" \
-d '{
"breed":"retriever",
"sub_breed":"golden"
}'

Example Response

{
  "message": "dog updated successfully"
}

------------------------------------------------------------

Delete Dog

DELETE /api/v1/dogs/{id}

Example curl

curl -X DELETE https://dogs-assignmnet-1.onrender.com/api/v1/dogs/1

Example Response

{
  "message": "dog deleted successfully"
}

------------------------------------------------------------

Health Check

GET /api/v1/health

Example curl

curl https://dogs-assignmnet-1.onrender.com/api/v1/health

Response

{
  "message": "!!! We are RunninGoo !!!"
}

------------------------------------------------------------

Persistence

All dog records are stored in PostgreSQL. Any changes made through the API are persisted in the database.

For example, if a dog record is deleted and the API is called again later, the deleted record will not appear because the data is permanently stored in the database.

------------------------------------------------------------

Deployment

The service is containerized with Docker and deployed using Render.

Live URL

https://dogs-assignmnet-1.onrender.com

------------------------------------------------------------

Notes

The application includes common middleware typically used in production services such as request logging, panic recovery, request timeout, and request ID generation.

The current implementation focuses on core CRUD functionality but the structure allows the service to be easily extended with features such as pagination, filtering, authentication, or API documentation.

------------------------------------------------------------

Author

Keerthi R
