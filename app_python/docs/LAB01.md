# Lab 01

## Framework Selection

I chose Flask because it is simple and enough for two endpoints. FastAPI has built-in API docs and async support, while Django has more features than this project needs.

| Framework | Footprint | Async support | OpenAPI/docs | Learning curve |
|---|---|---|---|---|
| Flask | Small | Async views with an extra dependency | Extensions needed | Low |
| FastAPI | Small | Native async | Built-in OpenAPI and interactive docs | Low to moderate |
| Django | Large | Async views supported | Third-party tools needed | Higher |

## Best Practices Applied

- Pinned Flask to keep the installed version consistent.
- Used environment variables to change settings without editing code.
- Used UTC timestamps to calculate uptime safely.
- Returned JSON and status codes so clients can handle errors.
- Added timestamps and levels to logs to make problems easier to find.
- Ignored generated files and the virtual environment to keep Git clean.

## API Documentation

### GET /

```bash
curl -s http://localhost:5050/ | jq .
```

The full response is in Testing Evidence below.

### GET /health

```bash
curl -s http://localhost:5050/health | jq .
```

```json
{
  "status": "ok",
  "timestamp": "2026-10-27T17:52:16.147013+00:00",
  "uptime_seconds": 4
}
```

## Testing Evidence

Ports 5000 and 8080 were busy, so I used 5050 and 8081. I started the first server from `app_python` with `HOST=127.0.0.1 PORT=5050 venv/bin/python app.py`.

```console
$ curl -s http://localhost:5050/ | jq .
{
  "endpoints": [
    {
      "description": "Service and host information",
      "method": "GET",
      "path": "/"
    },
    {
      "description": "Health check",
      "method": "GET",
      "path": "/health"
    }
  ],
  "request": {
    "client_ip": "127.0.0.1",
    "method": "GET",
    "path": "/",
    "user_agent": "curl/8.7.1"
  },
  "runtime": {
    "current_time": "2026-10-27T17:52:16.139100+00:00",
    "timezone": "UTC",
    "uptime_human": "4 seconds",
    "uptime_seconds": 4
  },
  "service": {
    "description": "Reports service and host information",
    "framework": "Flask",
    "name": "DevOps Info Service",
    "version": "1.0.0"
  },
  "system": {
    "architecture": "arm64",
    "cpu_count": 15,
    "hostname": "Igors-MacBook-Pro.local",
    "platform": "Darwin",
    "platform_version": "Darwin Kernel Version 25.6.0: Fri Jul 31 19:19:08 PDT 2026; root:xnu-12377.161.14~5/RELEASE_ARM64_T6050",
    "python_version": "3.14.7"
  }
}

$ curl -s http://localhost:5050/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T17:52:16.147013+00:00",
  "uptime_seconds": 4
}

$ curl -s -o /dev/null -w 'HTTP %{http_code}\n' http://localhost:5050/nope
HTTP 404
```

In another shell from `app_python`:

```console
$ PORT=8081 venv/bin/python app.py
2026-10-27 20:52:13,711 INFO Starting DevOps Info Service on 0.0.0.0:8081
 * Serving Flask app 'app'
 * Debug mode: off
 * Running on http://127.0.0.1:8081
```

```console
$ curl -s http://localhost:8081/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T17:52:16.163724+00:00",
  "uptime_seconds": 2
}
```

From the repository root:

```console
$ find app_python -maxdepth 2 -type f | sort
app_python/.gitignore
app_python/README.md
app_python/__pycache__/app.cpython-314.pyc
app_python/app.py
app_python/docs/LAB01.md
app_python/requirements.txt
app_python/tests/__init__.py
app_python/venv/.gitignore
app_python/venv/pyvenv.cfg
```

## Challenges & Solutions

The suggested ports were busy on my Mac. I checked which ports were in use and chose free ones to test the service.

## GitHub Community

IgorDuino

Stars show maintainers that people like their project and help others find it. Following developers helps me discover their projects and learn from their work. I starred both repositories, but my GitHub token could not follow users, so that step is still pending.
