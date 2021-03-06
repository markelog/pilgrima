FROM golang:1.11 AS builder

# Download and install the latest release of dep
# FIXME: do not use "ADD"
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/markelog/pilgrima
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

RUN mkdir /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app/run .

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app/run"]