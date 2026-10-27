# DevOps Info Service

```bash
cd app_python
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python app.py
```

For a different port:

```bash
PORT=8080 python app.py
```

- `GET /` — service and host info
- `GET /health` — health check

## Docker

Build and run from the repository root:

```bash
docker build -f app_python/Dockerfile.multi -t devops-info-service:lab02-multi ./app_python
docker run --rm -p 7071:5000 devops-info-service:lab02-multi
curl -s http://localhost:7071/health
```

Pull from GHCR:

```bash
docker pull ghcr.io/igorduino/devops-info-service:1.0.0
```
