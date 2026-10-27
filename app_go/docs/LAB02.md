# Lab 02 Bonus

The Go image uses a multi-stage build. `CGO_ENABLED=0` builds a static binary, so it does not need libc in the final image. I chose distroless because it has no shell or package manager and runs as a non-root user.

The app uses `uname` for system details when it is available. Distroless does not include it, so the Go app uses its runtime values and reports `unknown` for `platform_version`.

```console
$ docker images devops-info-service --format '{{.Tag}}\t{{.Size}}'
lab02-multi	218MB
lab02-go	18.4MB
$ file /tmp/lab02-go-binary
ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked
$ docker image inspect --format '{{.Config.User}}' devops-info-service:lab02-go
65532:65532
```

Trivy found 207 vulnerabilities in the Python image and 17 in the Go image: 3 Debian findings with unknown severity and 14 in the Go binary (4 high, 9 medium, 1 low). The Go image has fewer packages, but a small image is not automatically free of vulnerabilities. The binary was built with Go 1.25.14; Trivy reports fixes for some findings in newer Go releases. I would update the builder and rescan before deployment. Full output: [trivy-lab02.txt](trivy-lab02.txt).

Both endpoints worked in the container. The curl output is in the [Python lab report](../../app_python/docs/LAB02.md).
