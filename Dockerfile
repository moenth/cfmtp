# Debian go 1.7 image.
FROM golang:1.7

# Copy local package files to the container's workspace.
ADD . /go/src/github.com/moenth/cfmtp

# Build the app inside the container.
RUN go install github.com/moenth/cfmtp

# Run the cfmtp command when the container starts.
ENTRYPOINT /go/bin/cfmtp

# Expose app on :8080.
EXPOSE 8080
