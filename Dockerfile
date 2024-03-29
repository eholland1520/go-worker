# build stage
FROM golang:1.16.8-alpine3.14 AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
ADD . /src

RUN go get github.com/streadway/amqp
RUN go get github.com/tkanos/gonfig

RUN cd /src && go build -o go-worker

# final stage
FROM alpine
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=build-env /src/go-worker /app/
COPY --from=build-env /src/config.json /app/
COPY --from=build-env /src/wait-for-it.sh /app/
#ENTRYPOINT ./go-worker
CMD ["./go-worker"]