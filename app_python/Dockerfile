FROM python:3.14.7-slim
WORKDIR /app
RUN useradd  --uid 10001 --create-home --shell /bin/sh app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
USER 10001
EXPOSE 5000
HEALTHCHECK CMD ["python3", "-c", "import urllib.request; urllib.request.urlopen('http://localhost:5000/health', timeout=2)"]
CMD ["python3", "app.py"]
