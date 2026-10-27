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
