FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# Add the required build libraries
RUN apk update && apk add gcc librdkafka-dev zstd-libs libsasl lz4-dev libc-dev musl-dev 

# Copy and download dependency using go mod
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download

# Copy the code into the container
COPY ./src .

# Build the application
RUN go build -a -tags musl -ldflags="-extldflags=-static" -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

######################################
FROM scratch
LABEL AUTHOR="JULIAN BENSCH"
COPY /res/defaults.yml ./res/
COPY --from=builder /dist/main .
CMD ["./main"]