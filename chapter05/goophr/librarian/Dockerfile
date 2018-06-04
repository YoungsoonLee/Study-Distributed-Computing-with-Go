FROM golang:1.10

# In case of librarian, '/concierge' will be replaced with '/librarian'
ADD . /go/src/github.com/YoungsoonLee/distributed-go/chapter05/goophr/librarian

RUN go install github.com/YoungsoonLee/distributed-go/chapter05/goophr/librarian

ENTRYPOINT /go/bin/librarian

EXPOSE 9000