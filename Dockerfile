# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20.4 AS build-stage

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /diamonder


# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /app

COPY --from=build-stage /diamonder /diamonder

# Because it is multistage, we have to copy html and static files to container 
RUN mkdir views
COPY ./views ./views
RUN mkdir assets
COPY ./assets ./assets

EXPOSE 8080

ENTRYPOINT ["/diamonder"]