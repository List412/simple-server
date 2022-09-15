#FROM golang:1.18-alpine as app-builder
#WORKDIR /go/src/app
#COPY . .
#RUN apk add git
#ENV CGO_ENABLED=0
#RUN ls -lah
#RUN pwd
## RUN go test -cover ./internal/... \
#RUN go build -o /go/bin cmd/server
#RUN ls -lah bin
#
#FROM scratch
#COPY --from=app-builder /go/bin/server /server
#ENTRYPOINT ["/server"]

ARG  BUILDER_IMAGE=golang:buster
ARG  DISTROLESS_IMAGE=gcr.io/distroless/base
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on

# use modules
COPY ../go.mod .

ENV GO111MODULE=on
RUN go mod download && go mod verify

COPY . .

# Build the binary
RUN go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/server ./cmd/server

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}

# Copy our static executable
COPY --from=builder /go/bin/server /go/bin/server

# Run the hello binary.
ENTRYPOINT ["/go/bin/server"]