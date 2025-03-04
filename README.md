# Simple ToDo API in Golang

## Create a TODO:
```
curl -X POST -H "Content-Type: application/json" -d '{"title":"Learn Go","completed":false}' http://localhost:8080/todos
```

## Get All TODOs:
```
curl http://localhost:8080/todos
```

## Update a TODO:
```
curl -X PUT -H "Content-Type: application/json" -d '{"id":1,"title":"Learn GoLang","completed":true}' http://localhost:8080/todos
```

## Delete a TODO:
```
curl -X DELETE http://localhost:8080/todos/1
```
