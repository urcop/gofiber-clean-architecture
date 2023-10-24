# ✨ Go fiber template ✨
This is a basic template for the Fiber framework based on a clean architecture based 

## Start the application
- Copy .env.example to .env and enter your parameters `cp .env.examle .env`
- To run the application, use the `make run` commands to run your project locally or `make docker-build` to run your project in a docker container

## Swagger

http://localhost:8000/swagger/index.html

## Architecture

### /api
This is the directory for describing swagger files, to automatically generate a swagger use the command `make swagger`

### /build
The directory where the binary file of your project is located, to build the project use the command `make build`

### /cmd
The directory where the entry points to your applications are located

### /internal
The directory where all the logic of your project is located

### /pkg
The directory where all third-party packages are located

## Libraries used

- gofiber/fiber/v2 
- gobuffalo/envy ( To work with variables inside .env )
- gofiber/swagger ( To generate swagger files )
- swaggo/swag
- spf13/cobra ( To set flags to start the application )
- go.uber.org/zap ( logger )