FROM golang:1.9.1

# In case of librarian, '/concierge' will be replaced with '/librarian'
ADD . /go/src/github.com/YoungsoonLee/distributed-go/chapter05/goophr/concierge

WORKDIR /go/src/github.com/YoungsoonLee/distributed-go/chapter05/goophr/concierge

RUN go install github.com/YoungsoonLee/distributed-go/chapter05/goophr/concierge

ENTRYPOINT /go/bin/concierge

EXPOSE 9000