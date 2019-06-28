# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12.5
ARG GO_VERSION=1.12.5

# BUILD STAGE
FROM golang:${GO_VERSION}-alpine as builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Install git
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src
COPY . .

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
RUN go mod download

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# FINAL STAGE: the running container.
FROM scratch AS final

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled executable from the first stage.
COPY --from=builder /app /app
# Copy file .env for app
COPY --from=builder /src/config.toml .

EXPOSE 3000

ENTRYPOINT ["/app"]