# fuego cache
Fuego cache is a concurrent hashed key-value pair written 100% in Golang. Could run in 3 modes:
 - TPC process
 - HTTP server
 - CLI

### Introduction
Just need to deploy the fuego instance locally or in a cloud provider and connect with a tcp client or as a web server.
Need different "modes" if you need a TCP plain connection or a web server, add the environment variable MODE=tcp or
MODE=http as you wish, otherwise the std client will show up.

### Installation
TDB

### Modes
#### Operations in HTTP mode

1. Add a value

  URL: ```/fuego```
  
  Method: POST
  
  Body: 
  ```json
        {
          "key" : "someKey",
          "value" : "someValue"
        }
```
        
  Response: HTTP status = 201

2. Get a Value

  URL: ```/fuego/{key}```
  
  Method: GET
  
  Response:
  ```json
        {
          "key" : "someKey",
          "value" : "someValue"
        }
```
        
  status = 200
  
------------------------------------------------------------
