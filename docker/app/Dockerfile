# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Gustavo Gingins <gusgins@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go.mod, go.sum first and run go mod download to cache dependencias if files have not changed
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the working directory inside the container 
COPY . .

# Build go app as main from current the directory inside the container
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
#COPY --from=builder /app/docker/app/.env .
COPY --from=builder /app/config.yml .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
