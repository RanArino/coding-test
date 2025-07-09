# Stage 1: Build the test runner
FROM golang:1.22-alpine AS builder

RUN apk --no-cache add gcc musl-dev

WORKDIR /app

# Copy all Go modules and download dependencies
COPY tests/go.mod tests/go.sum ./
RUN go mod download

# Copy the source code
COPY tests/. .

# Build the test runner
RUN CGO_ENABLED=1 GOOS=linux go build -o /test_runner .

# Stage 2: Create the final runtime image
FROM golang:1.22-alpine

# Install dependencies
RUN apk --no-cache add gcc musl-dev sqlite dos2unix

# Copy the test runner binary from the builder stage
COPY --from=builder /test_runner /usr/local/bin/test_runner
COPY --from=builder /app/go.mod /app/go.mod
COPY --from=builder /app/go.sum /app/go.sum

# Copy the samples directory
COPY samples /app/samples

# Copy the test script and make it executable
COPY tests/run_tests.sh /app/run_tests.sh
RUN dos2unix /app/run_tests.sh
RUN chmod +x /app/run_tests.sh

# Set the working directory
WORKDIR /app

CMD ["./run_tests.sh"]
