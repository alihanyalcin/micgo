# micgo
**Generate microservice architecture based project with Go less than a second.**

micgo creates a starting point for your microservice-based application and you can only focus to develop your **REST APIs** needed by your business. 

The generated project has; 
* a **MongoDB** client, 
* **Dockerfile** for each microservices, 
* **Makefile** to build, run and dockerize the project, and
* **docker-compose** file.

# Usage

Download micgo:
```go
go get -u github.com/alihanyalcin/micgo
```
Generate your project:
```go
go run github.com/alihanyalcin/micgo generate testproject service1:12300 service2:12301
```
This command generates a project named **testproject**. **testproject** includes two services. One of the services is **service1** that serves on port **12300** and the other one is **service2** that serves on port **12301**.

Go to testproject directory and build it: 
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

To run the project with Docker Compose:
```sh
sudo make docker
```
