<div align="center">
      <h1><br/>To do list API</h1>
</div>

# Description
A simple Todo list API project with CRUD operations, authentication using JWT, and powered by Golang, Echo, Postgres, and implemented unit test on services and containerize using docker

# Features
This API is developed using Golang, echo web framework, GORM for database operations, and JWT for authentication.

# Tech Used
![Golang](https://img.shields.io/badge/golang-%23F7DF1E.svg?style=for-the-badge&logo=go&logoColor=black)
![Echo](https://img.shields.io/badge/echo-%2342D6AD.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![GORM](https://img.shields.io/badge/gorm-%2300f.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/jwt-%2300f.svg?style=for-the-badge&logo=jwt&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)


# Getting Started:
Before running the program, make sure you've installed the required dependencies.

- Install Golang: [Official Golang Installation Guide](https://golang.org/doc/install)
- Install Echo: `go get github.com/labstack/echo/v4`
- Install GORM: `go get gorm.io/gorm`
- Install POSTGRES driver for GORM: `go get gorm.io/driver/postgres"`
- Install JWT library: `go get github.com/golang-jwt/jwt/v5`

### Database setup:
Create your Postgres database and update the database configuration in the project.

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
#### Running Docker Container
```shell
docker run -e DB_HOST=host.docker.internal \
           -e DB_NAME=(YOUR_DB_NAME) \
           -e DB_PORT=(YOUR_DB_PORT) \
           -e DB_USER=(YOUR_DB_USER) \
           -e DB_PASSWORD= (YOUR_DB_PASSWORD)\
           -e SERVER_PORT=(YOUR_APP_PORT) \
           -e JWT_SECRET=(YOUR_JWT_PORT) \
           -p 8080:8080 \
           tiarajuliarsita/test-abc:v1
docker build -t tiarajuliarsita/test-abc:v1 .


```


