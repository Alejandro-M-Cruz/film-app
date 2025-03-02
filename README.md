# film-app
Simple API for a film app that allows users to register, login, and CRUD films. Authentication is done using JWT. 
All authenticated users can create and access films, but only the user that created a film can update or delete it.

## How to set up the project using Docker

### Requirements
- Docker
- Docker Compose

### Steps
1. Clone the repository: `git clone https://github.com/Alejandro-M-Cruz/film-app.git`
2. Create an .env file, copy all the variables in .env.example and fill the required values
3. Run the following command to build the Docker image: `docker compose up --build`
4. The project will be running on http://localhost:8000
5. Copy the OpenAPI documentation in `docs/api.yaml` to a [Swagger Editor](https://editor.swagger.io/)
6. Try out the API by registering a user, logging in, and using the JWT token to authenticate the other endpoints
   - You can also log in with the pre-registered user:
      - Username: `test`
      - Password: `password`


## How to set up the project locally

### Requirements
- Go 1.24
- SQLite 3
- Make

### Steps
1. Clone the repository: `git clone https://github.com/Alejandro-M-Cruz/film-app.git`
2. Install Go dependencies: `go mod download`
3. Create an .env file, copy all the variables in .env.example and fill the required values
4. Create and populate the SQLite database by running the following make command: `make db-fresh` 
5. Run the project: `go run main.go`
6. The project will be running on http://localhost:8000
7. Copy the OpenAPI documentation in `docs/api.yaml` to a [Swagger Editor](https://editor.swagger.io/)
8. Try out the API by registering a user, logging in, and using the JWT token to authenticate the other endpoints
    - You can also log in with the pre-registered user:
      - Username: `test`
      - Password: `password`
