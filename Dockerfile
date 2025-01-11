ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/templates ./templates
COPY --from=builder /usr/src/app/static ./static
COPY --from=builder /usr/src/app/markdown ./markdown
CMD ["run-app"]
