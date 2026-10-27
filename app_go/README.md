# DevOps Info Service in Go

Build and run with Go 1.21 or later:

```bash
cd app_go
CGO_ENABLED=0 go build -o devops-info .
./devops-info
```

Set `HOST` and `PORT` to change the address:

```bash
HOST=127.0.0.1 PORT=8080 ./devops-info
curl -s http://localhost:8080/ | jq .
curl -s http://localhost:8080/health | jq .
```

The JSON matches the Python service. The `system.python_version` field contains the Go version for this service.
