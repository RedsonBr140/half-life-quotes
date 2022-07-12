FROM golang:1.19-rc-alpine3.15
RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server .

CMD [ "/app/server" ]
