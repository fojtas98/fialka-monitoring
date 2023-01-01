FROM golang:alpine3.17
WORKDIR /app
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /src
CMD ["/src", ""]