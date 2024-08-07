FROM golang:1.22-alpine

RUN mkdir /app
WORKDIR /app

ADD gateway.go /app/
ADD go.mod /app/
ADD go.sum /app/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
RUN ls -l
CMD ["/main"]