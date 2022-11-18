FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o projects ./main.go

#CMD ["go", "run", "main.go"]
CMD ["./projects"]