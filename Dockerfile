# GO Repo base repo
FROM golang:1.20-alpine3.17 as builder
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Download all the dependencies
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN swag init

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o api .


#########################################################################
# GO Repo base repo
FROM alpine:3.17 as prod
RUN apk add tzdata
RUN addgroup -S api && adduser -S api -G api
RUN mkdir /app
WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

# Expose port 8080
EXPOSE 8080
USER api
# Run Executable
CMD ["./api"]
