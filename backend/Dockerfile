FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .

WORKDIR /app/cmd/server
RUN go mod download
RUN go build -o /app/main

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/main /app/main

# Copy static files
WORKDIR /app/images
COPY images /app/images

# Get back to root
WORKDIR /app

# Install ffmpeg for image processing
RUN apt-get -y update
RUN apt-get install -y ffmpeg

RUN chmod +x /app/main

EXPOSE 8080
CMD ["/app/main"]
