FROM golang:1.9.1

# In case of librarian, '/concierge' will be replaced with '/librarian'
ADD . /go/src/github.com/YoungsoonLee/distributed-go/chapter05/goophr/librarian

WORKDIR /go/src/github.com/YoungsoonLee/distributed-go/chapter05/goophr/librarian

RUN go install github.com/YoungsoonLee/distributed-go/chapter05/goophr/librarian

ENTRYPOINT /go/bin/librarian

EXPOSE 9000