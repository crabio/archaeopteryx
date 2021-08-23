# Build stage
FROM golang AS build-env
ADD . /src
ENV CGO_ENABLED=0
RUN cd /src && go build -o /app

# Production stage
FROM alpine:3
COPY --from=build-env /app /

# Create folder for logs
RUN mkdir /var/log/archaeropteryx

ENTRYPOINT ["/app"]
