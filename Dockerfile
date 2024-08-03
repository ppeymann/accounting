# Use the official Golang image as a build environment
FROM golang:alpine as build-env

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . ./

# Download the Go module dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /accounting ./cmd/accounting/main.go

# Use the official Alpine image as the base image for the final stage
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /

# Create necessary directories
RUN mkdir /data

# Add user and group
RUN addgroup --system accounting && adduser -S -s /bin/false -G accounting accounting

# Copy the built binary from the build stage
COPY --from=build-env /accounting /accounting

# Copy the config folder to the root directory of the container
COPY ./config/config.json /config/config.json

COPY ./schemas /schemas

COPY ./source/bank_import.json /source/bank_import.json

# Change ownership to the accounting user
RUN chown -R accounting:accounting /accounting
RUN chown -R accounting:accounting /data
RUN chown -R accounting:accounting /config
RUN chown -R accounting:accounting /schemas
RUN chown -R accounting:accounting /source

# Switch to the accounting user
USER accounting

# Expose the necessary port
EXPOSE 8080

# Set the entrypoint to the binary
ENTRYPOINT [ "/accounting" ]
