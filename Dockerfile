FROM golang:1.17.6-alpine3.15
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o bin/hack-iot -v .
EXPOSE 8880
CMD [ "/app/bin/hack-iot" ]