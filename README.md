# Food delivery

## overview
- Meaning of this repository: This is a reference point for everyone on the team to be able to work consistently
- Tech stack: Golang, Mysql, docker, docker-compose, GRPC, Socket.io, Jeager...
- Architecture: Based on the clean architecture
- Minimum understanding to work with this repo: 
  - Have a deep understanding of nodejs
  - Familiar with the clean architecture 
  - Have knowledge about mysql, sql

**Install `protoc`, related packages and generating swagger documents**

- Install `protoc`: http://google.github.io/proto-lens/installing-protoc.html
- Install related `protoc` packages (https://github.com/grpc-ecosystem/grpc-gateway)
- install `go`: https://golang.org/

```
  go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
  go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
  go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
  go get -u google.golang.org/protobuf/cmd/protoc-gen-go
```

**Start server**
- Project using Go 1.19.6 darwin/amd64
- Start your mysql by following: https://www.mysql.com/
- Start your gorm by following: https://gorm.io/
- Start your jaeger by following: https://www.jaegertracing.io/ (optional)
```bash
# development
$ go run main.go
```

 ## Application Structure

<img src="https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg" alt="clean-architecture-diagram-2.png" width="700">

##### Core: Entities
* Represent your domain object
* Apply only logic that is applicable in general to the whole entity (e.g. validating the format of an hostname)
* Plain objects: no frameworks, no annotations

##### Core: Use Cases
* Represent your business actions, it’s what you can do with the application. Expect one use case for each business action
* Pure business logic, plain typescrip (expect maybe some utils libraries like lodash, ramda)
* Define interfaces for the data that they need in order to apply some logic. One or more dataproviders will implement the interface, but the use case doesn’t know where the data is coming from
* The use case doesn't know who triggered it and how the results are going to be presented (e.g. could be on a web page, or returned as json, or simply logged, etc.)
* Throws business exceptions

##### Core: infrastructures
  - where we write the code to use external resource like: database, event bus...
  - `infrastructures/data-gateway`:
    * Retrieve and store data from and to a number of sources (database, network devices, file system, 3rd parties, etc.)
    * Implement the interfaces defined by the use case
    * Use whatever framework is most appropriate (they are going to be isolated here anyway)
    * Note: if using an ORM for database access, here you'd have another set of objects in order to represent the mapping to the tables (don't use the core entities as they might be very different)
    
##### Core: Transports
* Are ways to interact with the application, and typically involve a delivery mechanism (e.g. REST APIs, GRPC,  scheduled jobs, even-listener, other systems)
* Trigger a use case and convert the result to the appropriate format for the delivery mechanism
* The controller would trigger a use case

##### Configuration
* Wires everything together
* Frameworks (e.g. for dependency injection) are isolated here
* Has the "dirty details" like Main class, web server configuration, datasource configuration, etc.