import logging
import os
import platform
import socket
from datetime import datetime, timezone

from flask import Flask, jsonify, request


logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
app = Flask(__name__, static_folder=None)
START_TIME = datetime.now(timezone.utc)


def uptime_seconds():
    return int((datetime.now(timezone.utc) - START_TIME).total_seconds())


def uptime_human(seconds):
    units = ((86400, "day"), (3600, "hour"), (60, "minute"), (1, "second"))
    parts = []
    for size, name in units:
        count, seconds = divmod(seconds, size)
        if count:
            parts.append(f"{count} {name}{'s' if count != 1 else ''}")
    return ", ".join(parts) or "0 seconds"


@app.get("/")
def info():
    uptime = uptime_seconds()
    return jsonify({
        "service": {
            "name": "DevOps Info Service",
            "version": "1.0.0",
            "description": "Reports service and host information",
            "framework": "Flask",
        },
        "system": {
            "hostname": socket.gethostname(),
            "platform": platform.system(),
            "platform_version": platform.version(),
            "architecture": platform.machine(),
            "cpu_count": os.cpu_count(),
            "python_version": platform.python_version(),
        },
        "runtime": {
            "uptime_seconds": uptime,
            "uptime_human": uptime_human(uptime),
            "current_time": datetime.now(timezone.utc).isoformat(),
            "timezone": "UTC",
        },
        "request": {
            "client_ip": request.remote_addr,
            "user_agent": request.headers.get("User-Agent"),
            "method": request.method,
            "path": request.path,
        },
        "endpoints": [
            {"path": "/", "method": "GET", "description": "Service and host information"},
            {"path": "/health", "method": "GET", "description": "Health check"},
        ],
    })


@app.get("/health")
def health():
    return jsonify({
        "status": "ok",
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "uptime_seconds": uptime_seconds(),
    }), 200


@app.errorhandler(404)
def not_found(error):
    return jsonify({"error": "Not found"}), 404


@app.errorhandler(500)
def internal_error(error):
    return jsonify({"error": "Internal server error"}), 500


if __name__ == "__main__":
    host = os.getenv("HOST", "0.0.0.0")
    port = int(os.getenv("PORT", "5000"))
    debug = os.getenv("DEBUG", "False").lower() == "true"
    logging.info("Starting DevOps Info Service on %s:%s", host, port)
    app.run(host=host, port=port, debug=debug)
