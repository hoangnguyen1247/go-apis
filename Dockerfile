FROM golang:1.8

WORKDIR /go/src/app
COPY . .

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 8083
RUN go build
CMD ["go", "run", "main.go"]
