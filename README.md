# film-app
Simple API for a film app that allows users to register, login, and CRUD films. Authentication is done using JWT. 
All authenticated users can create and access films, but only the user that created a film can update or delete it.

## Requirements
- Go 1.24
- SQLite 3
- Make

## How to run the project locally
- Clone the repository: `git clone https://github.com/Alejandro-M-Cruz/film-app.git`
- Install Go dependencies: `go mod download`
- Create an .env file, copy all the variables in .env.example and fill the required values
- Create and populate the SQLite database by running the following make command: `make db-fresh` 
- Run the project: `go run main.go`
- The project will be running on http://localhost:8000
- Copy the OpenAPI documentation to the [Swagger Editor](https://editor.swagger.io/)

