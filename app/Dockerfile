FROM golang:1.11.4

RUN go get github.com/labstack/echo/...

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]
