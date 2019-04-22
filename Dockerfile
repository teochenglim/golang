FROM golang AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/teochenglim/golang
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM alpine
COPY --from=builder /app ./
# For nghttp2-dev, we need this respository.
RUN echo https://dl-cdn.alpinelinux.org/alpine/edge/testing >>/etc/apk/repositories

ENV CURL_VERSION 7.50.1

RUN apk add --update --no-cache ca-certificates curl

ENTRYPOINT ["./app"]

EXPOSE 3000
