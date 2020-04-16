# micGo
**Generate microservice architecture based project with Go less than a second.**

micGo creates a starting point for your microservice-based application and you can only focus to develop your **REST APIs** needed by your business.
 
<img src="micgo.png" align="right" width="300px" alt="micgo logo">

The generated project has; 
* **MongoDB** client for each services, 
* **Dockerfile** for each services, 
* **Makefile** to build, run and dockerize the project, and
* **docker-compose** file.

# Usage

Download **micGo**:
```go
go get -u github.com/alihanyalcin/micgo
```
Generate test project:
```go
go run github.com/alihanyalcin/micgo generate testproject service1:12300 service2:12301
```
This command generates a project named **testproject**. **testproject** has two services. One of the services is **service1** that serves on port **12300** and the other one is **service2** that serves on port **12301**.


>**NOTE:** If you want to generate **your project**, use the template:
>```go
>go run github.com/alihanyalcin/micgo generate <project_name> <service_name1>:<service_port1> ><service_name2>:<service_port2> ... <service_nameX>:<service_portX>
>```


Go to **testproject** directory and build it: 
```sh
cd testproject
make build
```
After build step is complete, run the project:
```sh
make run
```
OR
```sh
cd bin/
./launch.sh
```
**NOTE:** You must have **MongoDB** running on your system before launching the project.

Ping your services. You will get 'pong' message if everything is okay.
- http://localhost:12300/api/v1/ping for service1.
- http://localhost:12301/api/v1/ping for service2.

### To run the project with Docker Compose:
```sh
sudo make docker
sudo docker-compose up
```
**NOTE:** The Docker-compose file contains **MongoDB** and the microservices. So, if you have **MongoDB** running on your system, **stop** it. 

## Generated Project Structure
```
testproject/
├── bin
│   └── launch.sh                       # script to run all your services
├── cmd
│   ├── service1                        # main package of service1
│   │   ├── res                         # configuration files of service1
│   │   │   ├── docker
│   │   │   │   └── configuration.toml
│   │   │   └── configuration.toml
│   │   ├── Dockerfile                  # Docker file of service1
│   │   └── main.go                     # starting point of service1
│   └── service2                        # main package of service2
│       ├── res                         # configuration files of service2 
│       │   ├── docker
│       │   │   └── configuration.toml
│       │   └── configuration.toml
│       ├── Dockerfile                  # Docker file of service2
│       └── main.go                     # starting point of service2
├── internal                            # internal development area
│   ├── pkg                             # packages used in all services
│   │   ├── bootstrap                   # bootstrap package
│   │   │   ├── configuration
│   │   │   │   └── file.go
│   │   │   ├── container
│   │   │   │   ├── configuration.go
│   │   │   │   ├── database.go
│   │   │   │   └── logging.go
│   │   │   ├── handlers
│   │   │   │   ├── database
│   │   │   │   │   └── database.go
│   │   │   │   ├── httpserver
│   │   │   │   │   └── httpserver.go
│   │   │   │   └── message
│   │   │   │       └── message.go
│   │   │   ├── interfaces
│   │   │   │   ├── configuration.go
│   │   │   │   ├── database.go
│   │   │   │   └── handler.go
│   │   │   ├── logging
│   │   │   │   └── factory.go
│   │   │   ├── startup
│   │   │   │   └── timer.go
│   │   │   └── bootstrap.go
│   │   ├── config                      # configuration models package
│   │   │   └── types.go
│   │   ├── db                          # database package
│   │   │   ├── interfaces
│   │   │   │   └── db.go
│   │   │   ├── mongo
│   │   │   │   ├── client.go
│   │   │   │   └── test.go
│   │   │   └── db.go
│   │   ├── di                          # dependency injection package
│   │   │   ├── container.go
│   │   │   └── type.go
│   │   ├── logger                      # logger package
│   │   │   ├── log_entry.go
│   │   │   └── logger.go
│   │   ├── usage
│   │   │   └── usage.go
│   │   └── encoding.go
│   ├── service1                        # development area of service1
│   │   ├── config
│   │   │   └── config.go               # configuration of service1
│   │   ├── init.go                     # initialize service1
│   │   └── router.go                   # REST APIs of service1
│   ├── service2                        # development area of service2
│   │   ├── config
│   │   │   └── config.go               # configuration of service2
│   │   ├── init.go                     # initialize service2
│   │   └── router.go                   # REST APIs of service2
│   └── constants.go
├── docker-compose.yml                  # docker-compose file to deploy the project
├── go.mod                              # required modules
├── Makefile                            # makefile to build, run and dockerize
├── README.md
├── VERSION
└── version.go
```

