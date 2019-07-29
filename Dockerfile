FROM golang:1.12-alpine AS builder
LABEL maintainer="anton.fedyashov@gmail.com"

RUN apk add git 

ADD . /src
RUN export CGO_ENABLED=0 && cd /src/cmd/web && ls && go build -o /onlinekasse 

RUN wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest
RUN export CGO_ENABLED=0 && cd /src && golangci-lint run

FROM scratch
COPY --from=builder /onlinekasse /onlinekasse
CMD [ "/onlinekasse" ]

