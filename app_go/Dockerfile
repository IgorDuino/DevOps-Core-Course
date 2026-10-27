# syntax=docker/dockerfile:1
FROM golang:1.25 AS build
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /devops-info .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /devops-info /devops-info
USER 65532:65532
EXPOSE 5000
ENTRYPOINT ["/devops-info"]
