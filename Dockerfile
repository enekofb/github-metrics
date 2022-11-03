# Go build
FROM golang:1.19 AS go-build

WORKDIR /app
COPY . /app
COPY Makefile /app/
COPY go.* /app/
RUN go mod download
RUN make build

# Distroless
FROM gcr.io/distroless/base as runtime
COPY --from=go-build /app/metrics /metrics

ENTRYPOINT ["/metrics"]
