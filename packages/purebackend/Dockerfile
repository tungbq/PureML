FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git
RUN apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /tmp/purebackend

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/purebackend ./examples/prod/main.go

ARG PORT
ARG HOST_URL
ARG SCHEME
ARG DATABASE
ARG DATABASE_URL
ARG CGO_ENABLED
ARG PURE_ADMIN_AUTH_TOKEN_SECRET
ARG ADMIN_UUID
ARG ADMIN_EMAIL
ARG ADMIN_PASSWORD
ARG ADMIN_HANDLE
ARG ADMIN_ORG_HANDLE
ARG ADMIN_ORG_NAME
ARG PURE_R2_ACCOUNT_ID
ARG PURE_R2_BUCKET
ARG PURE_R2_ENDPOINT
ARG PURE_R2_ACCESS_KEY
ARG PURE_R2_SECRET
ARG PURE_S3_BUCKET
ARG PURE_S3_REGION
ARG PURE_S3_ENDPOINT
ARG PURE_S3_ACCESS_KEY
ARG PURE_S3_SECRET
ARG PURE_MAIL_HOST
ARG PURE_MAIL_PORT
ARG PURE_MAIL_USER
ARG PURE_MAIL_PASS
ARG PURE_SEARCH_HOST
ARG PURE_SEARCH_MASTER_KEY
ARG PURE_SEARCH_ADMIN_API_KEY
ARG PURE_SEARCH_API_KEY

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/purebackend/out/purebackend /app/purebackend

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/purebackend"]