############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 go build -o /go/bin/tester

# Create a "nobody" non-root user for the next image by crafting an /etc/passwd
# file that the next image can copy in. This is necessary since the next image
# is based on scratch, which doesn't have adduser, cat, echo, or even sh.
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/tester /go/bin/tester

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# # Copy the /etc_passwd file we created in the builder stage into /etc/passwd in
# # the target stage. This creates a new non-root user as a security best
# # practice.
COPY --from=builder /etc_passwd /etc/passwd

# # Run as the new non-root by default
USER nobody

# Run the hello binary.
ENTRYPOINT ["/go/bin/tester"]