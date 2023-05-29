
# Build the application from source
FROM golang:1.20 AS build-stage

# Install the build dependencies
RUN apt-get update && apt-get install -y \
    make \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN make build

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN make test

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/synthetic /synthetic

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/synthetic"]