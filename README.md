# fuego-cache
Fuego cache is a concurrent hashed key-value pair written 100% in Golang. Could run in 2 modes either as TPC process or with an HTTP server.

### Operations in HTTP mode

1. Add a value
  URL: ```/fuego```
  Method: POST
  Body: ```json
        {
          "key" : "someKey",
          "value" : "someValue"
        }```
  Response: 201

2. Get a Value
  URL: ```/fuego/{key}```
  Method: GET
  Response:```json
        {
          "key" : "someKey",
          "value" : "someValue"
        }```
  
