# micGo
**Generate microservice architecture based project with Go less than a second.**

micgo creates a starting point for your microservice-based application and you can only focus to develop your **REST APIs** needed by your business. 

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" align="right" width="430px" alt="logo">

The generated project has; 
* a **MongoDB** client, 
* **Dockerfile** for each microservices, 
* **Makefile** to build, run and dockerize the project, and
* **docker-compose** file.

# Usage

Download **micgo**:
```go
go get -u github.com/alihanyalcin/micgo
```
Generate test project:
```go
go run github.com/alihanyalcin/micgo generate testproject service1:12300 service2:12301
```
This command generates a project named **testproject**. **testproject** includes two services. One of the services is **service1** that serves on port **12300** and the other one is **service2** that serves on port **12301**.

------
**NOTE:** If you want to generate your project, use the template:
```go
go run github.com/alihanyalcin/micgo generate <project_name> <service_name1>:<service_port1> <service_name2>:<service_port2> ... <service_nameX>:<service_portX>
```
------

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
**NOTE:** You must have MongoDB running on your system before launching the project.

Ping your microservices. You will get 'pong' message if everything is okay.
- http://localhost:12300/api/v1/ping for service1.
- http://localhost:12301/api/v1/ping for service2.

To run the project with **Docker Compose**:
```sh
sudo make docker
sudo docker-compose up
```
**NOTE:** The Docker-compose file contains **MongoDB** and the microservices. So, if you have **MongoDB** running on your system, **stop** it. 

## Generated Project Structure

