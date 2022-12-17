FROM golang:1.18.3-alpine

RUN apk update && apk upgrade && apk add curl \
                          make \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN go mod tidy

RUN make build

EXPOSE 8801

ENTRYPOINT ["./main"]