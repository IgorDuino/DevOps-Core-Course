# Lab 02

from my first build. I later fixed the .dockerignore filename and rebuilt the images.

Dockerfile defendens:
1. slim to save storage and harder image so we take less CVEs to the project, pinned for Idempotency
2. uid >= 10000 - for security, avoid interception with system, potentially privileged users
3. I selected sh over bash for hardering.

dockerfile
```
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
```

dockerignore
```
.git
.venv
venv/
__pycache__/
node_modules/
.env
*.pyc
docs/
tests/
README.md
.idea/
.vscode/
```

build
```
 docker build -t devops-info-service:lab02 ./app_python
[+] Building 1.5s (11/11) FINISHED                                         docker:desktop-linux
 => [internal] load build definition from Dockerfile                                       0.0s
 => => transferring dockerfile: 400B                                                       0.0s
 => [internal] load metadata for docker.io/library/python:3.14.7-slim                      0.9s
 => [internal] load .dockerignore                                                          0.0s
 => => transferring context: 2B                                                            0.0s
 => [1/6] FROM docker.io/library/python:3.14.7-slim@sha256:51dafde81dbdb6ebde285137a295cf  0.0s
 => => resolve docker.io/library/python:3.14.7-slim@sha256:51dafde81dbdb6ebde285137a295cf  0.0s
 => [internal] load build context                                                          0.1s
 => => transferring context: 130.82kB                                                      0.1s
 => CACHED [2/6] WORKDIR /app                                                              0.0s
 => CACHED [3/6] RUN useradd  --uid 10001 --create-home --shell /bin/sh app                0.0s
 => CACHED [4/6] COPY requirements.txt .                                                   0.0s
 => CACHED [5/6] RUN pip install --no-cache-dir -r requirements.txt                        0.0s
 => [6/6] COPY . .                                                                         0.1s
 => exporting to image                                                                     0.4s
 => => exporting layers                                                                    0.3s
 => => exporting manifest sha256:9f8d6fcef46be17aebb84a8444ea43ccda7e431ba86b14326746ede9  0.0s
 => => exporting config sha256:eb14c9daa3ffbdaa21d4c50279d036610f9a3380fd0bf21471fc56d623  0.0s
 => => exporting attestation manifest sha256:11ed3fb3cf5909f10a7684771bf9243a4d8f888d8077  0.0s
 => => exporting manifest list sha256:1cae112b7151d366938b5b00635e9388272a59e11cf0940481a  0.0s
 => => naming to docker.io/library/devops-info-service:lab02                               0.0s
 => => unpacking to docker.io/library/devops-info-service:lab02                            0.1s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/l2gaasycge7ag6qld1qwylcq7
```

run and check
```
docker run -d --rm -p 7070:5000 --name lab02 devops-info-service:lab02

0b34653721e4c2a351f2ab23f3ef6829c3559f5e708ce64ddb50beb97e266b81
❯ curl -s http://localhost:7070/ | jq .
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
    "client_ip": "192.168.65.1",
    "method": "GET",
    "path": "/",
    "user_agent": "curl/8.7.1"
  },
  "runtime": {
    "current_time": "2026-10-27T19:10:33.913160+00:00",
    "timezone": "UTC",
    "uptime_human": "11 seconds",
    "uptime_seconds": 11
  },
  "service": {
    "description": "Reports service and host information",
    "framework": "Flask",
    "name": "DevOps Info Service",
    "version": "1.0.0"
  },
  "system": {
    "architecture": "aarch64",
    "cpu_count": 15,
    "hostname": "0b34653721e4",
    "platform": "Linux",
    "platform_version": "#1 SMP Fri May  1 14:35:41 UTC 2026",
    "python_version": "3.14.7"
  }
}
❯ curl -s http://localhost:7070/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T19:10:39.693036+00:00",
  "uptime_seconds": 17
}
```

non root
```
docker run --rm devops-info-service:lab02 id -u

10001
```
image
```
 docker images | grep lab
WARNING: This output is designed for human readability. For machine-readable output, please use --format.
devops-info-service:lab02                                                      1cae112b7151        252MB         54.8MB   U
```
history of image
```
docker history devops-info-service:lab02


IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
1cae112b7151   2 minutes ago   CMD ["python3" "app.py"]                        0B        buildkit.dockerfile.v0
<missing>      2 minutes ago   HEALTHCHECK &{["CMD" "python3" "-c" "import …   0B        buildkit.dockerfile.v0
<missing>      2 minutes ago   EXPOSE [5000/tcp]                               0B        buildkit.dockerfile.v0
<missing>      2 minutes ago   USER 10001                                      0B        buildkit.dockerfile.v0
<missing>      2 minutes ago   COPY . . # buildkit                             20.2MB    buildkit.dockerfile.v0
<missing>      6 minutes ago   RUN /bin/sh -c pip install --no-cache-dir -r…   17.1MB    buildkit.dockerfile.v0
<missing>      6 minutes ago   COPY requirements.txt . # buildkit              12.3kB    buildkit.dockerfile.v0
<missing>      6 minutes ago   RUN /bin/sh -c useradd  --uid 10001 --create…   69.6kB    buildkit.dockerfile.v0
<missing>      6 minutes ago   WORKDIR /app                                    8.19kB    buildkit.dockerfile.v0
<missing>      5 weeks ago     CMD ["python3"]                                 0B        buildkit.dockerfile.v0
<missing>      5 weeks ago     RUN /bin/sh -c set -eux;  for src in idle3 p…   16.4kB    buildkit.dockerfile.v0
<missing>      5 weeks ago     RUN /bin/sh -c set -eux;   savedAptMark="$(a…   45.2MB    buildkit.dockerfile.v0
<missing>      5 weeks ago     ENV PYTHON_SHA256=3b48dac8fb59f62eaa67ac83c1…   0B        buildkit.dockerfile.v0
<missing>      5 weeks ago     ENV PYTHON_VERSION=3.14.7                       0B        buildkit.dockerfile.v0
<missing>      5 weeks ago     RUN /bin/sh -c set -eux;  apt-get update;  a…   4.99MB    buildkit.dockerfile.v0
<missing>      5 weeks ago     ENV PATH=/usr/local/bin:/usr/local/sbin:/usr…   0B        buildkit.dockerfile.v0
<missing>      5 weeks ago     # debian.sh --arch 'arm64' out/ 'trixie' '@1…   110MB     debuerreotype 0.17
```

biggest - Bayes image, because containing all the system libraries, Python Interpreter itself. And everything else.

## Final results

### Dockerfile decisions

The single-stage Dockerfile and the final .dockerignore:

```dockerfile
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
```

```gitignore
.git
.venv
venv/
__pycache__/
node_modules/
.env
*.pyc
docs/
tests/
README.md
.idea/
.vscode/.env.*
. dockerignore
Dockerfile*
```

I used a pinned slim Python image to keep the runtime smaller and avoid unexpected version changes. The working directory is `/app`, and UID 10001 is outside the usual system account range. Requirements are copied first so code changes do not reinstall packages. `--no-cache-dir` keeps pip downloads out of the layer. The source stays owned by root and the app runs as UID 10001. The healthcheck uses Python because slim has no curl. Exec form sends signals directly to Python. The tag can still change upstream, so a digest would make the base more reproducible.

### Multi-stage results

```dockerfile
# syntax=docker/dockerfile:1
FROM python:3.14.7-slim AS builder
WORKDIR /builder
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

FROM python:3.14.7-slim AS runtime
RUN useradd --uid 10001 --no-create-home --shell /usr/sbin/nologin app && python -m pip uninstall -y pip
WORKDIR /app
COPY --from=builder /install /usr/local
COPY . .
USER 10001
EXPOSE 5000
HEALTHCHECK CMD ["python3", "-c", "import urllib.request; urllib.request.urlopen('http://localhost:5000/health', timeout=2)"]
CMD ["python3", "app.py"]
```

```
$ docker images devops-info-service --format '{{.Tag}}\t{{.Size}}'
lab02-multi	218MB
lab02	226MB
lab02-go	18.4MB
```

The multi-stage image is 8 MB smaller here. Flask has no compiled extensions that need a build toolchain, so the builder does not remove a large toolchain from this build. The runtime removes pip; Docker still keeps lower layers, so image size is not the same as files visible at runtime. The Debian base layer is the largest in the single-stage history at 110 MB.

### Trivy scan

Trivy 0.76.1 scanned the final multi-stage image on 2026-10-27. Full output: [trivy-lab02.txt](trivy-lab02.txt).

| Severity | Count |
|---|---:|
| CRITICAL | 0 |
| HIGH | 48 |
| MEDIUM | 88 |
| LOW | 67 |
| UNKNOWN | 4 |
| Total | 207 |

```
$ trivy image --skip-db-update --severity HIGH,CRITICAL --exit-code 1 --format table --output /tmp/lab02-python-gate.txt devops-info-service:lab02-multi
$ echo $?
1
```

One finding is `libpcre2-8-0` 10.46-1~deb13u2, CVE-2026-103111. Trivy reports an out-of-bounds write when a crafted regular expression is used with certain JIT APIs. The scanner lists 10.46-1~deb13u3 as fixed. I would rebuild from a base with that Debian update, then scan again. [Finding](https://avd.aquasec.com/nvd/cve-2026-103111).

### Build / run / push evidence

```
$ docker build --progress=plain -t devops-info-service:lab02 app_python

#11 exporting to image
#11 exporting layers done
#11 exporting manifest sha256:e606f804512b888b71221cc2f818e78d5b5a7d7ec79df222f24a6bb647eec5c6 done
#11 exporting config sha256:a4934966ff81de6805b94a141cb566e3f15be17c7f84551ac4ed19ed6cf7a25a done
#11 exporting attestation manifest sha256:14b25176e2b1a9ef460c6f1a60805d346f99a02188f417b057746981da0c6b87 done
#11 exporting manifest list sha256:f0df688c35597f4cb4b5e01181085ec0376d4f6745273ce6e1f94a063ff94c5f done
#11 naming to docker.io/library/devops-info-service:lab02 done
#11 unpacking to docker.io/library/devops-info-service:lab02 done
#11 DONE 0.1s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/otd3wtb5zlhrdn0ky2qrhh6mp

$ docker build --progress=plain -f app_python/Dockerfile.multi -t devops-info-service:lab02-multi app_python
#15 exporting to image
#15 exporting layers 0.0s done
#15 exporting manifest sha256:cdaa9b3250cda3bec4697fb0ffb5cabe32b33b31ea02efa20cad3a01562e3d3f
#15 exporting manifest sha256:cdaa9b3250cda3bec4697fb0ffb5cabe32b33b31ea02efa20cad3a01562e3d3f done
#15 exporting config sha256:57566068bbd5b306d60546653d91d5fcf8d468b19c5f4f3d903b93fb94c04a97 done
#15 exporting attestation manifest sha256:212e0ce5a0c3e9a28f608f57853b0db584d2ddb53cbacba1448eb84916376fdf 0.0s done
#15 exporting manifest list sha256:a7ad600aed16c5336558164546150552eb3b674e6fc0d506bff83603da8728f3 done
#15 naming to docker.io/library/devops-info-service:lab02-multi done
#15 unpacking to docker.io/library/devops-info-service:lab02-multi 0.1s done
#15 DONE 0.2s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/6dy3s7pq2d4afa2qnv2cnt9zj
```

The ports below were assigned by Docker for the test.

```
$ docker run -d --rm -p 127.0.0.1::5000 devops-info-service:lab02
e7dbd340fb4c7936fd3de23339e634b7c07ae69cf9e68b45d1e26272ee6f843e
$ docker port e7dbd340fb4c7936fd3de23339e634b7c07ae69cf9e68b45d1e26272ee6f843e 5000/tcp
127.0.0.1:55296
$ curl -s http://127.0.0.1:55296/ | jq .
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
    "client_ip": "172.17.0.1",
    "method": "GET",
    "path": "/",
    "user_agent": "curl/8.7.1"
  },
  "runtime": {
    "current_time": "2026-10-27T19:51:39.956110+00:00",
    "timezone": "UTC",
    "uptime_human": "0 seconds",
    "uptime_seconds": 0
  },
  "service": {
    "description": "Reports service and host information",
    "framework": "Flask",
    "name": "DevOps Info Service",
    "version": "1.0.0"
  },
  "system": {
    "architecture": "aarch64",
    "cpu_count": 15,
    "hostname": "e7dbd340fb4c",
    "platform": "Linux",
    "platform_version": "#1 SMP Fri May  1 14:35:41 UTC 2026",
    "python_version": "3.14.7"
  }
}
$ curl -s http://127.0.0.1:55296/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T19:51:39.974620+00:00",
  "uptime_seconds": 0
}
$ docker run -d --rm -p 127.0.0.1::5000 devops-info-service:lab02-multi
2fbba7389e3902aea65e28aa18bac15794bf8cefaa28b8115077cc5778e42618
$ docker port 2fbba7389e3902aea65e28aa18bac15794bf8cefaa28b8115077cc5778e42618 5000/tcp
127.0.0.1:55311
$ curl -s http://127.0.0.1:55311/ | jq .
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
    "client_ip": "172.17.0.1",
    "method": "GET",
    "path": "/",
    "user_agent": "curl/8.7.1"
  },
  "runtime": {
    "current_time": "2026-10-27T19:51:41.475098+00:00",
    "timezone": "UTC",
    "uptime_human": "0 seconds",
    "uptime_seconds": 0
  },
  "service": {
    "description": "Reports service and host information",
    "framework": "Flask",
    "name": "DevOps Info Service",
    "version": "1.0.0"
  },
  "system": {
    "architecture": "aarch64",
    "cpu_count": 15,
    "hostname": "2fbba7389e39",
    "platform": "Linux",
    "platform_version": "#1 SMP Fri May  1 14:35:41 UTC 2026",
    "python_version": "3.14.7"
  }
}
$ curl -s http://127.0.0.1:55311/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T19:51:41.488364+00:00",
  "uptime_seconds": 0
}
$ docker run -d --rm -p 127.0.0.1::5000 devops-info-service:lab02-go
ac81fafb1f9d41113ff3a2cd58cf0203a68583571e54994a5c47d9bf4845ef04
$ docker port ac81fafb1f9d41113ff3a2cd58cf0203a68583571e54994a5c47d9bf4845ef04 5000/tcp
127.0.0.1:55331
$ curl -s http://127.0.0.1:55331/ | jq .
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
    "client_ip": "172.17.0.1",
    "method": "GET",
    "path": "/",
    "user_agent": "curl/8.7.1"
  },
  "runtime": {
    "current_time": "2026-10-27T19:51:42.875382799Z",
    "timezone": "UTC",
    "uptime_human": "0 seconds",
    "uptime_seconds": 0
  },
  "service": {
    "description": "Reports service and host information",
    "framework": "net/http",
    "name": "DevOps Info Service",
    "version": "1.0.0"
  },
  "system": {
    "architecture": "arm64",
    "cpu_count": 15,
    "hostname": "ac81fafb1f9d",
    "platform": "linux",
    "platform_version": "unknown",
    "python_version": "go1.25.14"
  }
}
$ curl -s http://127.0.0.1:55331/health | jq .
{
  "status": "ok",
  "timestamp": "2026-10-27T19:51:42.891402383Z",
  "uptime_seconds": 0
}

$ docker run --rm devops-info-service:lab02 id -u
10001
$ docker run --rm devops-info-service:lab02-multi id -u
10001
$ docker run --rm --entrypoint python devops-info-service:lab02-multi -m pip --version
/usr/local/bin/python: No module named pip
```

Registry: https://github.com/users/IgorDuino/packages/container/package/devops-info-service

```console
$ docker push ghcr.io/igorduino/devops-info-service:lab2
The push refers to repository [ghcr.io/igorduino/devops-info-service]
f08687e05e6a: Layer already exists
bd36565c0fde: Layer already exists
9b598a1e0170: Layer already exists
3e1b82b51359: Layer already exists
a34e28b955fd: Layer already exists
f370b8d18785: Layer already exists
d7cb81352efc: Layer already exists
e7234800df46: Layer already exists
e54df8b4e257: Already exists
lab2: digest: sha256:a7ad600aed16c5336558164546150552eb3b674e6fc0d506bff83603da8728f3 size: 856
$ docker logout ghcr.io
Removing login credentials for ghcr.io
$ docker pull ghcr.io/igorduino/devops-info-service:lab2
lab2: Pulling from igorduino/devops-info-service
Digest: sha256:a7ad600aed16c5336558164546150552eb3b674e6fc0d506bff83603da8728f3
Status: Image is up to date for ghcr.io/igorduino/devops-info-service:lab2
ghcr.io/igorduino/devops-info-service:lab2
```

The pull above used a local image cache. I also fetched the arm64 manifest, config and every layer from GHCR without registry credentials and verified their SHA-256 digests. The published image supports linux/arm64.

I also tagged the same image with a SemVer tag:

```console
$ docker tag devops-info-service:lab02-multi ghcr.io/igorduino/devops-info-service:1.0.0
$ docker push ghcr.io/igorduino/devops-info-service:1.0.0
1.0.0: digest: sha256:a7ad600aed16c5336558164546150552eb3b674e6fc0d506bff83603da8728f3 size: 856
```

I pulled it using a temporary Docker configuration with no registry credentials. The image was already cached locally:

```console
$ docker --config <empty-config-dir> pull ghcr.io/igorduino/devops-info-service:1.0.0
1.0.0: Pulling from igorduino/devops-info-service
Digest: sha256:a7ad600aed16c5336558164546150552eb3b674e6fc0d506bff83603da8728f3
Status: Image is up to date for ghcr.io/igorduino/devops-info-service:1.0.0
ghcr.io/igorduino/devops-info-service:1.0.0
```

The anonymous manifest for `1.0.0` has the same digest as `lab2`, whose layers I checked above.

[Docker run commands](../README.md).

### Challenges & Solutions

My first `.dockerignore` had a space in its name, so Docker sent the venv in the build context. I added the correct filename and rebuilt. The application layer dropped from 20.2 MB to a few KB.
