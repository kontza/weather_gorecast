FROM golang:1-alpine as builder

ARG PACKAGE_SOURCE
ARG PACKAGE_APP

# Basic requirements.
RUN apk update \
 && apk --no-cache add git \
 && go get -u $PACKAGE_SOURCE

# Set our workdir to our current service in the gopath
WORKDIR $GOPATH/src/$PACKAGE_SOURCE
COPY . .

# Here we're pulling in godep, which is a dependency manager tool,
# we're going to use dep instead of go get, to get around a few
# quirks in how go get works with sub-packages.
RUN go get -u github.com/golang/dep/cmd/dep

# Run `ensure`, which will pull in all
# of the dependencies within this directory.
RUN dep ensure

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -a .

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

ARG PACKAGE_SOURCE
ARG PACKAGE_APP

# Security related package, good to have.
RUN apk update && apk --no-cache add ca-certificates

# Same as before, create a directory for our app.
RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/$PACKAGE_SOURCE/$PACKAGE_APP service

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependencies and
# run time libraries.
CMD ["/app/service"]
