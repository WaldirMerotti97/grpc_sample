This project aims to show how to implement the 4 kinds of service method availble in gRPC.

Unary RPC: single request and single responde, as we are used to REST.
Server Streaming RPC: client sends a request to the server and receives a stream to read a sequence of messages.
Client Streaming RPC: client sends a sequence of messages to the server, the server reads and process all of them, returning the result at once.
Bidirectional Streaming RPC: Both client and server send messages to each other in a read-write stream.

How to run the project:

There are a few ways you can run/test this project. The following way was tested and works pretty well , using wsl2 along with VSCode.

Steps:

- Install the Go language installed in your machine

You can follow the step-by-step from this page: https://dev.to/deadwin19/how-to-install-golang-on-wslwsl2-2880

- Use an IDE of preference, I used VSCode to develop the project along with wsl2.

- Download the necessary dependencies, run the following command:

go mod tidy

- Running the code:

go run cmd/grpcServer/main.go 

- Calling the rpcs from the service definition:

I'm using evans, and grpc http client to call the services created and attached to the grpc server. There, you'll find how to use it!

https://github.com/ktr0731/evans

- Install sqlite3 in wsl to create a database and a table to store the application data.

First, update the packages using the command sudo apt update

Then, install SQLite3 using the command sudo apt install sqlite3

Create the database named db.sqlite with the following command sqlite3 db.sqlite


