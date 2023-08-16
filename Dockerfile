# Use the official Golang image as the base for building
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the app source code to the container
COPY . .


COPY prefectures-stations.json /app/prefectures-stations.json
COPY prefectures.json /app/prefectures.json
COPY cities.json /app/cities.json

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o api-app .

RUN apt-get update 
RUN apt-get install -y zip
