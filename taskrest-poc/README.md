# Task Rest POC

## Prerequisites
- Go 1.24.4
  
## How to run

```bash
go mod tidy
go run main.go
```

## API Usage
### Add a task
```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"title":"Learn Go"}' \
     http://localhost:9000/tasks
```

### List all tasks
```bash
curl http://localhost:8080/tasks
```

### Mark task as done
```bash
curl -X PUT http://localhost:8080/tasks/1
```

### Delete a task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```