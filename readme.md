<div align="center">
      <h1><br/>To do list API</h1>
</div>

# Description
A simple Todo list API project with CRUD operations, authentication using JWT, and powered by Golang, Echo, Postgres, and implemented unit test on services and containerize using docker

# Tech Used
![Golang](https://img.shields.io/badge/golang-%23F7DF1E.svg?style=for-the-badge&logo=go&logoColor=black)
![Echo](https://img.shields.io/badge/echo-%2342D6AD.svg?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/postgresql-%2300f.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![GORM](https://img.shields.io/badge/gorm-%2300f.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)



# Getting Started:
Before running the program, make sure you've installed the required dependencies.

- Install Golang: [Official Golang Installation Guide](https://golang.org/doc/install)
- Install Echo: `go get github.com/labstack/echo/v4`
- Install GORM: `go get gorm.io/gorm`
- Install POSTGRES driver for GORM: `go get gorm.io/driver/postgres"`
- Install JWT library: `go get github.com/golang-jwt/jwt/v5`

# Database setup:
Create your Postgres database and update the database configuration in the project.

# Env Configuration
For environment-specific configurations, please rename the example.env file to .env and update it with your project-specific environment variables. Make sure to set the appropriate values for configurations

### Run the program:
```shell
go run main.go
```

### Docker Set Up
If you prefer running the application in Docker containers, follow these steps:

#### Pull Docker image
Pull the Docker image from your Docker repository:
```shell
docker pull tiarajuliarsita/test-abc:v1
```
or if you want to build yourself, you can run:

```shell
docker pull (username)/(name of image):(version)
```
#### Running Docker Container
```shell
docker run -e DB_HOST=host.docker.internal \
           -e DB_NAME=(YOUR_DB_NAME) \
           -e DB_PORT=(YOUR_DB_PORT) \
           -e DB_USER=(YOUR_DB_USER) \
           -e DB_PASSWORD= (YOUR_DB_PASSWORD)\
           -e SERVER_PORT=(YOUR_APP_PORT) \
           -e JWT_SECRET=(YOUR_JWT_PORT) \
           -p (port_host):(port_container) \
           tiarajuliarsita/test-abc:v1
```

# Unit Test
Using testify for testing code and generate mock with Mockery.
- Install Mockery
```shell
go install github.com/vektra/mockery/v2@v2.38.0
```
- Generate mock
  before generate, make sure you are on the right path
```shell
mockery -all  
```

- Run Unit test
run with coverage
```shell
go test -v -cover   
```

# API Collection
Download the API collection from the following link to get started

[API collection](https://drive.google.com/drive/folders/1_XBbrNNol9Dhvu8jRRvU9lLO5yGv4zAG)




