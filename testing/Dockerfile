FROM golang:1 AS builder

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12
ENV GO111MODULE=on

WORKDIR "$GOPATH/src/github.com/Nekroze/ishmael"

COPY ./go.* ./

RUN go get -d -v ./...

COPY . .

RUN golangci-lint run --deadline '2m' --enable-all --disable dupl,lll,gochecknoglobals,gochecknoinits \
 && go test ./... \
 && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o /go/bin/ishmael


FROM nekroze/containaruba:alpine AS test

COPY --from=builder /go/bin/ishmael /usr/bin/ishmael
